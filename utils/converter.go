package utils

import (
	"errors"
	"reflect"
	"strings"
	"time"

	"github.com/emicklei/pgtalk/convert"
	"github.com/jackc/pgx/v5/pgtype"
)

// Function to create a field map for easier access
func createFieldMap(val reflect.Value) map[string]reflect.Value {
	fieldMap := make(map[string]reflect.Value)
	for i := 0; i < val.NumField(); i++ {
		field := val.Type().Field(i)
		fieldMap[strings.ToLower(field.Name)] = val.Field(i)
	}
	return fieldMap
}

// PrimitiveMappingParser function definition goes here...

func MappingParser(sender interface{}, receiver interface{}) error {
	sentVal := reflect.ValueOf(sender)
	receivedVal := reflect.ValueOf(receiver)

	// Check if both parameters are pointers to structs
	if sentVal.Kind() != reflect.Ptr || receivedVal.Kind() != reflect.Ptr {
		return errors.New("both parameters must be pointers to structs")
	}

	// Dereference the pointers to access the underlying structs
	sentVal = sentVal.Elem()
	receivedVal = receivedVal.Elem()

	// Check if both values are structs
	if sentVal.Kind() != reflect.Struct || receivedVal.Kind() != reflect.Struct {
		return errors.New("both parameters must be structs")
	}

	// Create a map for the receiver struct's fields for O(1) lookups
	receiverFieldMap := createFieldMap(receivedVal)

	// Iterate over the fields of the sender struct
	for i := 0; i < sentVal.NumField(); i++ {
		senderField := sentVal.Type().Field(i)
		senderValue := sentVal.Field(i)

		// Perform a constant-time lookup in the map
		if receiverField, ok := receiverFieldMap[strings.ToLower(senderField.Name)]; ok && receiverField.CanSet() {
			// Handle conversions
			switch {
			case receiverField.Type() == senderValue.Type():
				// Direct assignment if types match
				receiverField.Set(senderValue)

			case receiverField.Type() == reflect.TypeOf(pgtype.UUID{}): // pgtype.UUID
				switch senderValue.Kind() {
				case reflect.String: // string to pgtype.UUID
					var uuid pgtype.UUID
					err := uuid.Scan(senderValue.String())
					if err == nil {
						receiverField.Set(reflect.ValueOf(uuid))
					}
				}

			case receiverField.Type() == reflect.TypeOf(pgtype.Text{}): // pgtype.Text
				switch senderValue.Kind() {
				case reflect.String: // string to pgtype.Text
					var text pgtype.Text
					err := text.Scan(senderValue.String())
					if err == nil {
						receiverField.Set(reflect.ValueOf(text))
					}
				}

			case receiverField.Type() == reflect.TypeOf(pgtype.Timestamptz{}): // pgtype.Timestamptz
				switch senderValue.Kind() {
				case reflect.Struct: // time.Time to pgtype.Timestamptz
					var timestamp pgtype.Timestamptz
					err := timestamp.Scan(senderValue.Interface().(time.Time))
					if err == nil {
						receiverField.Set(reflect.ValueOf(timestamp))
					}
				}
			

			// Allowing conversion from pgtype.UUID to string
			case senderValue.Type() == reflect.TypeOf(pgtype.UUID{}):// Allowing conversion to string
				if uuidValue, ok := senderValue.Interface().(pgtype.UUID); ok {
					receiverField.Set(reflect.ValueOf(convert.UUIDToString(uuidValue)))
				}

			// Allowing conversion from pgtype.Text to string
			case senderValue.Type() == reflect.TypeOf(pgtype.Text{}): // Allowing conversion to string
				
				if textValue, ok := senderValue.Interface().(pgtype.Text); ok {
					receiverField.Set(reflect.ValueOf(textValue.String))
				}

			// Allowing conversion from pgtype.Timestamptz to time.Time
			case senderValue.Type() == reflect.TypeOf(pgtype.Timestamptz{}): // Allowing conversion to time.Time
				if timeValue, ok := senderValue.Interface().(pgtype.Timestamptz); ok {
					receiverField.Set(reflect.ValueOf(timeValue.Time))
				}
			}
		}
	}

	return nil
}