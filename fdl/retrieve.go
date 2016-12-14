package fdl

// Property retrieves property data
func Property(property string) PropertyItem {
	p := PropertyItem{
		Property: property,
		Type:     DataType,
		IsPrefix: false,
	}
	return p
}

// Set retrieves property set data
func Set(property string) PropertyItem {
	p := PropertyItem{
		Property: property,
		Type:     SetType,
		IsPrefix: false,
	}
	return p
}

// Counter retrieves property counter
func Counter(property string) PropertyItem {
	p := PropertyItem{
		Property: property,
		Type:     CounterType,
		IsPrefix: false,
	}
	return p
}

// Meta retrieves property meta data
func Meta(property string) PropertyItem {
	p := PropertyItem{
		Property: property,
		Type:     MetaType,
		IsPrefix: false,
	}
	return p
}

// PropertiesWithPrefix retrieves property data prefixed with given key
func PropertiesWithPrefix(property string) PropertyItem {
	p := PropertyItem{
		Property: property,
		Type:     DataType,
		IsPrefix: true,
	}
	return p
}

// CountersWithPrefix retrieves property counters prefixed with given key
func CountersWithPrefix(property string) PropertyItem {
	p := PropertyItem{
		Property: property,
		Type:     CounterType,
		IsPrefix: true,
	}
	return p
}

// MetaWithPrefix retrieves property meta data prefixed with given key
func MetaWithPrefix(property string) PropertyItem {
	p := PropertyItem{
		Property: property,
		Type:     MetaType,
		IsPrefix: true,
	}
	return p
}
