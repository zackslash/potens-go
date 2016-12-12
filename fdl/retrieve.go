package fdl

// Property retrieves property data
func Property(property string) PropertyItem {
	p := PropertyItem{
		Property: property,
		Type:     DataType,
	}
	return p
}

// Set retrieves property set data
func Set(property string) PropertyItem {
	p := PropertyItem{
		Property: property,
		Type:     SetType,
	}
	return p
}

// Counter retrieves property counter
func Counter(property string) PropertyItem {
	p := PropertyItem{
		Property: property,
		Type:     CounterType,
	}
	return p
}

// Meta retrieves property meta data
func Meta(property string) PropertyItem {
	p := PropertyItem{
		Property: property,
		Type:     MetaType,
	}
	return p
}
