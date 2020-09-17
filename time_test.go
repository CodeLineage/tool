package tool

import(
	"testing"
)

func TestTimeUnix(t *testing.T){
	t.Log(TimeUnix())
}

func TestDateTime(t *testing.T){
    t.Log(DateTime(""))
    t.Log(DateTime("2006-01-02"))
}

func TestFormatStringToTime(t *testing.T){
    datetime := "2020-09-17 10:48:00"
    rs ,err := FormatStringToTime("", datetime)
    if err != nil{
        t.Fatal(err)
    }
    t.Log(rs)
}
