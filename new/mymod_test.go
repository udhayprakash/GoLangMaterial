package hello
 
import "testing"
 
func TestHello(t *testing.T) {
    want := "My module"
    if got := GetModName(); got != want {
        t.Errorf("Hello() = %q, want %q", got, want)
    }
}