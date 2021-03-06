package structures

import (
	"testing"
	"fmt"
)

func TestElement(t *testing.T) {
	f := NewCurrencyField(0, "")
	fieldElementRaw := element{nil, f.field}
	fieldElementConstructed := NewElement(f.field, nil)
	cases := []struct {
		got  element
		want element
	}{
		{fieldElementConstructed, fieldElementRaw},
	}
	for _, c := range cases {
		if !c.got.Equals(c.want) {
			t.Errorf("got == %q, want %q", c.got, c.want)
		}
	}
}

func TestElement_ToString(t *testing.T) {
	fieldElementWithValue := NewElement(NewCurrencyField(4, "123").field,
		nil)

	blockWithValue := NewElement(NewCurrencyField(3, "123").field,
		[]element{
			NewElement(NewCurrencyField(4, "123").field,
				nil),
			NewElement(NewCurrencyField(4, "123").field,
				nil),
		})

	blockWithBlockWithValue := NewElement(NewCurrencyField(3, "123").field,
		[]element{
			NewElement(NewCurrencyField(4, "123").field,
				[]element{
					NewElement(NewCurrencyField(4, "100").field,
						nil),
					NewElement(NewCurrencyField(4, "100").field,
						nil),
				}),

			NewElement(NewCurrencyField(4, "123").field,
				nil)},
		)

	cases := []struct {
		got  string
		want string
	}{
		{fieldElementWithValue.ToString(), "0123"},
		{blockWithValue.ToString(), "12301230123"},
		{blockWithBlockWithValue.ToString(), "1230123010001000123"},
	}
	for _, c := range cases {
		if c.got != c.want {
			t.Errorf("got == %q, want %q", c.got, c.want)
		}
	}
}

func TestElement_Parse(t *testing.T) {
	got := NewElement(NewCurrencyField(4, "").field,
		[]element{
			NewElement(NewCurrencyField(4, "").field,
				[]element{
					NewElement(NewCurrencyField(4, "").field,
						nil),
					NewElement(NewCurrencyField(4, "").field,
						[]element{
							NewElement(NewCurrencyField(4, "3000").field,
								nil),
							NewElement(NewCurrencyField(4, "4000").field,
								nil),
						}),
				}),

			NewElement(NewCurrencyField(4, "").field,
				nil)},
		)

	want := NewElement(NewCurrencyField(4, "1000").field,
		[]element{
			NewElement(NewCurrencyField(4, "2000").field,
				[]element{
					NewElement(NewCurrencyField(4, "3000").field,
						nil),
					NewElement(NewCurrencyField(4, "4000").field,
						[]element{
							NewElement(NewCurrencyField(4, "3000").field,
								nil),
							NewElement(NewCurrencyField(4, "4000").field,
								nil),
						}),
				}),

			NewElement(NewCurrencyField(4, "5000").field,
				nil)},
		)
	input := "1000200030004000300040005000"


	got.Parse(input)
	fmt.Printf("\n got: %s\n", got.ToString())
	fmt.Printf("want: %s\n", want.ToString())

	if !got.Equals(want){
		t.Errorf("\n got == %s\nwant == %s", got, want)
	}

	failInput := "9999999999999999999999999999"
	got.Parse(failInput)
	fmt.Printf(" got: %s\n", got.ToString())
	fmt.Printf("want: %s\n", want.ToString())

	if got.Equals(want){
		t.Errorf("\nValues are equal, should not be.\n got == %s\nwant == %s", got, want)
	}
}
