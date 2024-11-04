package usecases

import (
	"context"
	"fmt"
	"log"

	"github.com/ThanawatPtd/SAProject/config"
	"github.com/ThanawatPtd/SAProject/domain/repositories"
	"github.com/sendgrid/sendgrid-go"
	"github.com/sendgrid/sendgrid-go/helpers/mail"
)

type EmailUseCase interface {
	GetExpiredTransactionThisWeek(ctx context.Context) error
	SendReceipt(ctx context.Context, encodedFile string, filename string,id string) (error)
}

type EmailService struct {
	tranRepo repositories.TransactionRepository
	key      string
}


func ProvideEmailService(tranRepo repositories.TransactionRepository, config *config.Config) EmailUseCase {
	return &EmailService{
		tranRepo: tranRepo,
		key:      config.SendGridSecret,
	}
}

func (e *EmailService) SendReceipt(ctx context.Context, encodedFile string, filename string, id string) error {
	transaction, err := e.tranRepo.GetUserVehicleTransactionByID(ctx, id)
	if err != nil {
		return err
	}

	from := mail.NewEmail("RUMUE Service", "lerdphipat.k@ku.th")
	subject := "Receipt from RUMUE"
	to := mail.NewEmail("K."+transaction.User.Fname+" "+transaction.User.Lname, "lerdphipat.k@ku.th")

	plainTextContent := "Thank you for choosing RUMUE! Your car insurance details and policy duration are provided below. We're here to keep you covered!"

	emailContent := fmt.Sprintf(`
		<!DOCTYPE html>
		<html lang="en">
		<head>
			<style>
				body { font-family: Arial, sans-serif; color: #333; margin: 20px; }
				.container { width: 100%%; max-width: 600px; margin: 0 auto; }
				.header { text-align: center; padding: 20px 0; background-color: #f4f4f4; }
				.details { padding: 20px; background-color: #ffffff; margin-top: 10px; border-radius: 8px; box-shadow: 0 2px 4px rgba(0,0,0,0.1); }
				.section { margin-bottom: 20px; }
				.footer { text-align: center; font-size: 12px; color: #777; margin-top: 20px; }
				.cta { display: inline-block; padding: 10px 20px; margin-top: 20px; background-color: #4a90e2; color: #ffffff; text-decoration: none; border-radius: 5px; }
			</style>
		</head>
		<body>
			<div class="container">
				<div class="header">
					<h1>Receipt for Your Insurance Purchase</h1>
					<p>Hello %s, thank you for purchasing your car insurance with RUMUE.</p>
				</div>
				<div class="details">
					<div class="section payment-info">
						<h2>Payment Details</h2>
						<p><strong>Insurance Type:</strong> %s</p>
						<p><strong>Status:</strong> %s</p>
						<p><strong>Amount Paid:</strong> %.2f THB</p>
						<p><strong>Transaction ID:</strong> %s</p>
					</div>
					<div class="section car-info">
						<h2>Car Information</h2>
						<p><strong>Make & Model:</strong> %s %s</p>
						<p><strong>Year:</strong> %s</p>
						<p><strong>Color:</strong> %s</p>
						<p><strong>Seating Capacity:</strong> %d</p>
						<p><strong>Engine Number:</strong> %s</p>
						<p><strong>Chassis Number:</strong> %s</p>
					</div>
					<div class="section thank-you">
						<h2>Thank You for Your Purchase!</h2>
						<p>Your policy is now active. You can manage your policy details, view receipts, and make changes by logging into your RUMUE account.</p>
						<a href="http://localhost:3000/history" class="cta">View My Account</a>
					</div>
				</div>
				<div class="footer">
					<p>If you have any questions, contact us at <a href="mailto:lerdphipat.k@ku.th">lerdphipat.k@ku.th</a> or call 092-272-0521.</p>
					<p>Thank you for choosing RUMUE!</p>
				</div>
			</div>
		</body>
		</html>`,
		transaction.User.Fname,
		transaction.Transaction.InsuranceType,
		transaction.Transaction.Status,
		transaction.Transaction.Price,
		transaction.Transaction.ID,
		transaction.Vehicle.Brand,
		transaction.Vehicle.Model,
		transaction.Vehicle.ModelYear,
		transaction.Vehicle.VehicleColor,
		transaction.Vehicle.SeatingCapacity,
		transaction.Vehicle.EngineNumber,
		transaction.Vehicle.ChasisNumber,
	)

	message := mail.NewSingleEmail(from, subject, to, plainTextContent, emailContent)

	attachment := mail.NewAttachment()
	attachment.SetContent(encodedFile)
	attachment.SetType("application/pdf")
	attachment.SetFilename(filename)
	attachment.SetDisposition("attachment")
	message.AddAttachment(attachment)


	client := sendgrid.NewSendClient(e.key)
	response, err := client.Send(message)
	if err != nil {
		log.Println("Error sending email:", err)
		return err
	}

	log.Printf("Email sent! Status Code: %d\nResponse Body: %s\nResponse Headers: %v\n", response.StatusCode, response.Body, response.Headers)
	return nil
}


func (e *EmailService) GetExpiredTransactionThisWeek(ctx context.Context) error {
	transaction, err := e.tranRepo.GetExpiredTransactionThisWeek(ctx)
	if err != nil {
		return err
	}

	for _, val := range transaction {
		from := mail.NewEmail("RUMUE Service", "lerdphipat.k@ku.th")
		subject := "Your Car Insurance Renewal Notice"
		to := mail.NewEmail("K."+val.User.Fname, val.User.Email)

		// For testing in development env only

		// transaction := []UserVehicleTransaction{
		// 	{
		// 		User: User{
		// 			ID:          "user123",
		// 			Fname:       "John",
		// 			Lname:       "Doe",
		// 			Email:       "john.doe@example.com",
		// 			PhoneNumber: "092-272-0521",
		// 		},
		// 		Transaction: Transaction{
		// 			ID:            "trans123",
		// 			Price:         1500.00,
		// 			InsuranceType: "Comprehensive",
		// 			Status:        "Active",
		// 		},
		// 		Vehicle: Vehicle{
		// 			ID:              "vehicle123",
		// 			Brand:           "Toyota",
		// 			Model:           "Corolla",
		// 			ModelYear:       "2021",
		// 			VehicleColor:    "Red",
		// 			SeatingCapacity: 5,
		// 			EngineNumber:    "ENG123456",
		// 			ChasisNumber:    "CHAS123456",
		// 		},
		// 	},
		// }

		plainTextContent := "Thank you for choosing RUMUE! Your car insurance details and policy duration are provided below. We're here to keep you covered!"
		emailContent := fmt.Sprintf(`<!DOCTYPE html>
							<html lang="en">
							<head>
								<style>
									body { font-family: Arial, sans-serif; color: #333; margin: 20px; }
									.container { width: 100%%; max-width: 600px; margin: 0 auto; }
									.header { text-align: center; padding: 20px 0; background-color: #f4f4f4; }
									.details { padding: 20px; background-color: #ffffff; margin-top: 10px; border-radius: 8px; box-shadow: 0 2px 4px rgba(0,0,0,0.1); }
									.section { margin-bottom: 20px; }
									.car-info, .policy-info, .renewal-info { margin-bottom: 14px; }
									.footer { text-align: center; font-size: 12px; color: #777; margin-top: 20px; }
									strong { color: #4a90e2; }
									.alert { color: #d9534f; font-weight: bold; }
									.cta { display: inline-block; padding: 10px 20px; margin-top: 20px; background-color: #4a90e2; color: #ffffff; text-decoration: none; border-radius: 5px; }
									.warning { color: #d9534f; font-weight: bold; margin-top: 10px; }
								</style>
							</head>
							<body>
								<div class="container">
									<div class="header">
										<h1>Important Notice: Insurance Expiration Soon</h1>
										<p>Hello %s, Your car insurance policy with RUMUE will expire in one week.</p>
									</div>

									<div class="details">
										<div class="section policy-info">
											<h2>Policy Details</h2>
											<p><strong>Insurance Type:</strong> %s</p>
											<p><strong>Status:</strong> %s</p>
											<p><strong>Last Price:</strong> %.2f</p>
										</div>

										<div class="section car-info">
											<h2>Car Information</h2>
											<p><strong>Make & Model:</strong> %s %s</p>
											<p><strong>Year:</strong> %s</p>
											<p><strong>Color:</strong> %s</p>
											<p><strong>Seating Capacity:</strong> %d</p>
											<p><strong>Engine Number:</strong> %s</p>
											<p><strong>Chassis Number:</strong> %s</p>
										</div>

										<div class="section renewal-info">
											<h2 class="alert">Renew Your Insurance</h2>
											<p>To continue your coverage seamlessly, please renew your policy before it expires. RUMUE makes this easy—just log in to your account, where your car's details are already saved, and complete the renewal in a few clicks.</p>
											<p class="warning">**Important:** Please recheck your vehicle information before making the purchase to ensure everything is accurate.</p>
											<a href="https://your-rumue-website.com/login" class="cta">Renew Insurance Now</a>
										</div>
									</div>

									<div class="footer">
										<p>If you have any questions, contact us at <a href="mailto:lerdphipat.k@ku.th">lerdphipat.k@ku.th</a> or call 092-272-0521.</p>
										<p>Thank you for driving safely with RUMUE!</p>
									</div>
								</div>
							</body>
							</html>
							`,
			val.User.Fname,
			val.Transaction.InsuranceType,
			val.Transaction.Status,
			val.Transaction.Price,
			val.Vehicle.Brand,
			val.Vehicle.Model,
			val.Vehicle.ModelYear,
			val.Vehicle.VehicleColor,
			val.Vehicle.SeatingCapacity,
			val.Vehicle.EngineNumber,
			val.Vehicle.ChasisNumber,
		)

		message := mail.NewSingleEmail(from, subject, to, plainTextContent, emailContent)

		client := sendgrid.NewSendClient(e.key)
		response, err := client.Send(message)
		if err != nil {
			log.Println("Error sending email:", err)
			return err
		} else {
			fmt.Println("Email sent! Status Code:", response.StatusCode)
			fmt.Println("Response Body:", response.Body)
			fmt.Println("Response Headers:", response.Headers)
		}
	}

	return nil
}
