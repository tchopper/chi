package chi

import "testing"

func TestURLParamMethods(t *testing.T) {
	c := Context{}

	val := c.URLParam("key")
	if val != "" {
		t.Errorf("URLParam function failed - val should be ''")
	}

	c.SetURLParam("key", "value")

	val = c.URLParam("key")
	if val != "value" {
		t.Errorf("URLParam function failed - val should be 'value'")
	}

	c.SetURLParam("key", "valueOverride")
	val = c.URLParam("key")
	if val != "valueOverride" {
		t.Errorf("URLParam function failed - val should be 'valueOverride'")
	}

}
