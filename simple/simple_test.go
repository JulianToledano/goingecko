package simple_test
import (
	"testing"

	"github.com/JulianToledano/goingecko"
)

func TestSimplePrice(t *testing.T) {
	cgClient := goingecko.NewClient(nil)

	r, _ := cgClient.SimplePrice("bitcoin,ethereum", "btc,eth", true, true,true,true)
	if r == nil {
		t.Errorf("Error")
	}

}

func TestSimpleSupportedVsCurrency(t *testing.T) {
	cgClient := goingecko.NewClient(nil)

	supVcurr, _ := cgClient.SimpleSupportedVsCurrency()
	if supVcurr == nil {
		t.Errorf("Error")
	}
}
