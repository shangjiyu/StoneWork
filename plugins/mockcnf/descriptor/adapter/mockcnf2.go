// Code generated by adapter-generator. DO NOT EDIT.

package adapter

import (
	. "go.ligato.io/vpp-agent/v3/plugins/kvscheduler/api"
	"go.pantheon.tech/stonework/proto/mockcnf"
	"google.golang.org/protobuf/proto"
)

////////// type-safe key-value pair with metadata //////////

type MockCnf2KVWithMetadata struct {
	Key      string
	Value    *mockcnf.MockCnf2
	Metadata interface{}
	Origin   ValueOrigin
}

////////// type-safe Descriptor structure //////////

type MockCnf2Descriptor struct {
	Name                 string
	KeySelector          KeySelector
	ValueTypeName        string
	KeyLabel             func(key string) string
	ValueComparator      func(key string, oldValue, newValue *mockcnf.MockCnf2) bool
	NBKeyPrefix          string
	WithMetadata         bool
	MetadataMapFactory   MetadataMapFactory
	Validate             func(key string, value *mockcnf.MockCnf2) error
	Create               func(key string, value *mockcnf.MockCnf2) (metadata interface{}, err error)
	Delete               func(key string, value *mockcnf.MockCnf2, metadata interface{}) error
	Update               func(key string, oldValue, newValue *mockcnf.MockCnf2, oldMetadata interface{}) (newMetadata interface{}, err error)
	UpdateWithRecreate   func(key string, oldValue, newValue *mockcnf.MockCnf2, metadata interface{}) bool
	Retrieve             func(correlate []MockCnf2KVWithMetadata) ([]MockCnf2KVWithMetadata, error)
	IsRetriableFailure   func(err error) bool
	DerivedValues        func(key string, value *mockcnf.MockCnf2) []KeyValuePair
	Dependencies         func(key string, value *mockcnf.MockCnf2) []Dependency
	RetrieveDependencies []string /* descriptor name */
}

////////// Descriptor adapter //////////

type MockCnf2DescriptorAdapter struct {
	descriptor *MockCnf2Descriptor
}

func NewMockCnf2Descriptor(typedDescriptor *MockCnf2Descriptor) *KVDescriptor {
	adapter := &MockCnf2DescriptorAdapter{descriptor: typedDescriptor}
	descriptor := &KVDescriptor{
		Name:                 typedDescriptor.Name,
		KeySelector:          typedDescriptor.KeySelector,
		ValueTypeName:        typedDescriptor.ValueTypeName,
		KeyLabel:             typedDescriptor.KeyLabel,
		NBKeyPrefix:          typedDescriptor.NBKeyPrefix,
		WithMetadata:         typedDescriptor.WithMetadata,
		MetadataMapFactory:   typedDescriptor.MetadataMapFactory,
		IsRetriableFailure:   typedDescriptor.IsRetriableFailure,
		RetrieveDependencies: typedDescriptor.RetrieveDependencies,
	}
	if typedDescriptor.ValueComparator != nil {
		descriptor.ValueComparator = adapter.ValueComparator
	}
	if typedDescriptor.Validate != nil {
		descriptor.Validate = adapter.Validate
	}
	if typedDescriptor.Create != nil {
		descriptor.Create = adapter.Create
	}
	if typedDescriptor.Delete != nil {
		descriptor.Delete = adapter.Delete
	}
	if typedDescriptor.Update != nil {
		descriptor.Update = adapter.Update
	}
	if typedDescriptor.UpdateWithRecreate != nil {
		descriptor.UpdateWithRecreate = adapter.UpdateWithRecreate
	}
	if typedDescriptor.Retrieve != nil {
		descriptor.Retrieve = adapter.Retrieve
	}
	if typedDescriptor.Dependencies != nil {
		descriptor.Dependencies = adapter.Dependencies
	}
	if typedDescriptor.DerivedValues != nil {
		descriptor.DerivedValues = adapter.DerivedValues
	}
	return descriptor
}

func (da *MockCnf2DescriptorAdapter) ValueComparator(key string, oldValue, newValue proto.Message) bool {
	typedOldValue, err1 := castMockCnf2Value(key, oldValue)
	typedNewValue, err2 := castMockCnf2Value(key, newValue)
	if err1 != nil || err2 != nil {
		return false
	}
	return da.descriptor.ValueComparator(key, typedOldValue, typedNewValue)
}

func (da *MockCnf2DescriptorAdapter) Validate(key string, value proto.Message) (err error) {
	typedValue, err := castMockCnf2Value(key, value)
	if err != nil {
		return err
	}
	return da.descriptor.Validate(key, typedValue)
}

func (da *MockCnf2DescriptorAdapter) Create(key string, value proto.Message) (metadata Metadata, err error) {
	typedValue, err := castMockCnf2Value(key, value)
	if err != nil {
		return nil, err
	}
	return da.descriptor.Create(key, typedValue)
}

func (da *MockCnf2DescriptorAdapter) Update(key string, oldValue, newValue proto.Message, oldMetadata Metadata) (newMetadata Metadata, err error) {
	oldTypedValue, err := castMockCnf2Value(key, oldValue)
	if err != nil {
		return nil, err
	}
	newTypedValue, err := castMockCnf2Value(key, newValue)
	if err != nil {
		return nil, err
	}
	typedOldMetadata, err := castMockCnf2Metadata(key, oldMetadata)
	if err != nil {
		return nil, err
	}
	return da.descriptor.Update(key, oldTypedValue, newTypedValue, typedOldMetadata)
}

func (da *MockCnf2DescriptorAdapter) Delete(key string, value proto.Message, metadata Metadata) error {
	typedValue, err := castMockCnf2Value(key, value)
	if err != nil {
		return err
	}
	typedMetadata, err := castMockCnf2Metadata(key, metadata)
	if err != nil {
		return err
	}
	return da.descriptor.Delete(key, typedValue, typedMetadata)
}

func (da *MockCnf2DescriptorAdapter) UpdateWithRecreate(key string, oldValue, newValue proto.Message, metadata Metadata) bool {
	oldTypedValue, err := castMockCnf2Value(key, oldValue)
	if err != nil {
		return true
	}
	newTypedValue, err := castMockCnf2Value(key, newValue)
	if err != nil {
		return true
	}
	typedMetadata, err := castMockCnf2Metadata(key, metadata)
	if err != nil {
		return true
	}
	return da.descriptor.UpdateWithRecreate(key, oldTypedValue, newTypedValue, typedMetadata)
}

func (da *MockCnf2DescriptorAdapter) Retrieve(correlate []KVWithMetadata) ([]KVWithMetadata, error) {
	var correlateWithType []MockCnf2KVWithMetadata
	for _, kvpair := range correlate {
		typedValue, err := castMockCnf2Value(kvpair.Key, kvpair.Value)
		if err != nil {
			continue
		}
		typedMetadata, err := castMockCnf2Metadata(kvpair.Key, kvpair.Metadata)
		if err != nil {
			continue
		}
		correlateWithType = append(correlateWithType,
			MockCnf2KVWithMetadata{
				Key:      kvpair.Key,
				Value:    typedValue,
				Metadata: typedMetadata,
				Origin:   kvpair.Origin,
			})
	}

	typedValues, err := da.descriptor.Retrieve(correlateWithType)
	if err != nil {
		return nil, err
	}
	var values []KVWithMetadata
	for _, typedKVWithMetadata := range typedValues {
		kvWithMetadata := KVWithMetadata{
			Key:      typedKVWithMetadata.Key,
			Metadata: typedKVWithMetadata.Metadata,
			Origin:   typedKVWithMetadata.Origin,
		}
		kvWithMetadata.Value = typedKVWithMetadata.Value
		values = append(values, kvWithMetadata)
	}
	return values, err
}

func (da *MockCnf2DescriptorAdapter) DerivedValues(key string, value proto.Message) []KeyValuePair {
	typedValue, err := castMockCnf2Value(key, value)
	if err != nil {
		return nil
	}
	return da.descriptor.DerivedValues(key, typedValue)
}

func (da *MockCnf2DescriptorAdapter) Dependencies(key string, value proto.Message) []Dependency {
	typedValue, err := castMockCnf2Value(key, value)
	if err != nil {
		return nil
	}
	return da.descriptor.Dependencies(key, typedValue)
}

////////// Helper methods //////////

func castMockCnf2Value(key string, value proto.Message) (*mockcnf.MockCnf2, error) {
	typedValue, ok := value.(*mockcnf.MockCnf2)
	if !ok {
		return nil, ErrInvalidValueType(key, value)
	}
	return typedValue, nil
}

func castMockCnf2Metadata(key string, metadata Metadata) (interface{}, error) {
	if metadata == nil {
		return nil, nil
	}
	typedMetadata, ok := metadata.(interface{})
	if !ok {
		return nil, ErrInvalidMetadataType(key)
	}
	return typedMetadata, nil
}
