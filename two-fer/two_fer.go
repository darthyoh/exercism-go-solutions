// Simple package exporting a simple two fer function
package twofer

// ShareWith return a twofer string
func ShareWith(name string) string {

	if name == "" {
		name = "you"
	}
	return "One for " + name + ", one for me."
}
