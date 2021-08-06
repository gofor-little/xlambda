package xlambda

import (
	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"github.com/gofor-little/xerror"
)

// UnmarshalDynamoDBEventAttributeValues unmarshals attributeValues into v. For v to work with attributevalue.UnmarshalMap
// it must be a non-nil reference to a struct.
//
// CAUTION: If one of the types of the v interface{} parameter contains a slice of interfaces where an item is of type int
// it will be converted into a float64. I believe this is either a bug or limitation of the Go reflect package rather
// than an AWS packages. I haven't look into it enough to be sure though.
//
// This helper function is required because the Lambda events package uses different types to the DynamoDB package.
// This became even more of a pain the the v2 of the Go AWS SDK. See the following for more information.
// https://github.com/aws/aws-lambda-go/issues/58
// https://github.com/aws/aws-sdk-go-v2/issues/1124
func UnmarshalDynamoDBEventAttributeValues(attributeValues map[string]events.DynamoDBAttributeValue, v interface{}) error {
	attributes := make(map[string]types.AttributeValue, len(attributeValues))
	var err error

	for k, v := range attributeValues {
		attributes[k], err = fromDynamoDBAttributeValue(v)
		if err != nil {
			return xerror.Wrap("failed to parse events.DynamoDBAttributeValue of type map into types.AttributeValue", err)
		}
	}

	return attributevalue.UnmarshalMap(attributes, v)
}

// fromDynamoDBAttributeValue will convert an events.DynamoDBAttributeValue into a types.AttributeValue.
func fromDynamoDBAttributeValue(eventValue events.DynamoDBAttributeValue) (types.AttributeValue, error) {
	switch eventValue.DataType() {
	case events.DataTypeBinary:
		return &types.AttributeValueMemberB{Value: eventValue.Binary()}, nil
	case events.DataTypeBoolean:
		return &types.AttributeValueMemberBOOL{Value: eventValue.Boolean()}, nil
	case events.DataTypeBinarySet:
		return &types.AttributeValueMemberBS{Value: eventValue.BinarySet()}, nil
	case events.DataTypeList:
		values, err := fromDynamoDBAttributeValueList(eventValue.List())
		if err != nil {
			return nil, xerror.Wrap("failed to parse events.DynamoDBAttributeValue of type list into types.AttributeValue", err)
		}

		return &types.AttributeValueMemberL{Value: values}, nil
	case events.DataTypeMap:
		values, err := fromDynamoDBAttributeValueMap(eventValue.Map())
		if err != nil {
			return nil, xerror.Wrap("failed to parse events.DynamoDBAttributeValue of type map into types.AttributeValue", err)
		}

		return &types.AttributeValueMemberM{Value: values}, nil
	case events.DataTypeNumber:
		return &types.AttributeValueMemberN{Value: eventValue.Number()}, nil
	case events.DataTypeNumberSet:
		return &types.AttributeValueMemberNS{Value: eventValue.NumberSet()}, nil
	case events.DataTypeNull:
		return &types.AttributeValueMemberNULL{Value: eventValue.IsNull()}, nil
	case events.DataTypeString:
		return &types.AttributeValueMemberS{Value: eventValue.String()}, nil
	case events.DataTypeStringSet:
		return &types.AttributeValueMemberSS{Value: eventValue.StringSet()}, nil
	default:
		// Fallthrough.
	}

	return nil, xerror.Newf("failed to parse data type: %v", eventValue.DataType())
}

// fromDynamoDBAttributeValueList will convert a []events.DynamoDBAttributeValue into a []types.AttributeValue.
func fromDynamoDBAttributeValueList(eventValues []events.DynamoDBAttributeValue) ([]types.AttributeValue, error) {
	typeValues := make([]types.AttributeValue, len(eventValues))
	var err error

	for i := 0; i < len(eventValues); i++ {
		typeValues[i], err = fromDynamoDBAttributeValue(eventValues[i])
		if err != nil {
			return nil, xerror.Wrap("failed to parse types.AttributeValue into events.DynamoDBAttributeValue", err)
		}
	}

	return typeValues, nil
}

// fromDynamoDBAttributeValueMap will convert a map[string]events.DynamoDBAttributeValue into a map[string]types.AttributeValue.
func fromDynamoDBAttributeValueMap(eventValues map[string]events.DynamoDBAttributeValue) (map[string]types.AttributeValue, error) {
	typeValues := make(map[string]types.AttributeValue, len(eventValues))
	var err error

	for k, v := range eventValues {
		typeValues[k], err = fromDynamoDBAttributeValue(v)
		if err != nil {
			return nil, xerror.Wrap("failed to parse types.AttributeValue into events.DynamoDBAttributeValue", err)
		}
	}

	return typeValues, nil
}
