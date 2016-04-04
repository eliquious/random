package main

import (
    "crypto/rand"
    "encoding/base64"
    "encoding/hex"
    "fmt"
    "github.com/Pallinder/go-randomdata"
    "github.com/alecthomas/kingpin"
    "io"
    mrand "math/rand"
    "time"
)

var (
    sillyNameCommand       = kingpin.Command("silly", "Generate a random silly name.")
    firstNameCommand       = kingpin.Command("first", "Generate a random first name.")
    firstNameGender        = firstNameCommand.Arg("GENDER", "Gender of generated name").Required().Enum("male", "female")
    lastNameCommand        = kingpin.Command("last", "Generate a random last name.")
    fullNameCommand        = kingpin.Command("full", "Generate a random full name.")
    fullNameGender         = fullNameCommand.Arg("GENDER", "Gender of generated name").Required().Enum("male", "female")
    emailCommand           = kingpin.Command("email", "Generate a random email.")
    countryCommand         = kingpin.Command("country", "Generate a random country.")
    twoLetterCountryFlag   = countryCommand.Flag("two", "Limits the country name to 2 letters").Bool()
    threeLetterCountryFlag = countryCommand.Flag("three", "Limits the country name to 3 letters").Bool()
    currencyCommand        = kingpin.Command("currency", "Generates a random currency.")
    cityCommand            = kingpin.Command("city", "Generates a random city.")
    stateCommand           = kingpin.Command("state", "Generates a random state.")
    shortStateFlag         = stateCommand.Flag("short", "Limits the state name to 2 letters").Bool()
    streetCommand          = kingpin.Command("street", "Generates a random street.")
    addressCommand         = kingpin.Command("address", "Generates a random address.")
    numberCommand          = kingpin.Command("number", "Generates a random number.")
    startNumberArg         = numberCommand.Arg("START", "Determines the start of the random number").Required().Int()
    stopNumberArg          = numberCommand.Arg("STOP", "Determines the stop of the random number").Required().Int()
    decimalCommand         = kingpin.Command("decimal", "Generates a random decimal.")
    startDecimalArg        = decimalCommand.Arg("START", "Determines the start of the random decimal").Required().Int()
    stopDecimalArg         = decimalCommand.Arg("STOP", "Determines the stop of the random decimal").Required().Int()
    precDecimalArg         = decimalCommand.Arg("PREC", "Determines the precision of the random decimal").Required().Int()
    boolCommand            = kingpin.Command("bool", "Generates a random bool.")
    paragraphCommand       = kingpin.Command("paragraph", "Generates a random paragraph.")
    postalCommand          = kingpin.Command("postal", "Generates a random postal.")
    postalArg              = postalCommand.Arg("COUNTRY", "Determines the country of the postal code. (2 letters)").Required().String()
    ipCommand              = kingpin.Command("ip", "Generates a random ipv4.")
    dayCommand             = kingpin.Command("day", "Generates a random day.")
    hexCommand             = kingpin.Command("hex", "Generates a random hex.")
    hexLengthArg           = hexCommand.Arg("LENGTH", "Determines the length of the random data in bytes").Default("16").Int()
    base64Command          = kingpin.Command("base64", "Generates a random base64.")
    base64LengthArg        = base64Command.Arg("LENGTH", "Determines the length of the random data in bytes").Default("16").Int()
    stringCommand          = kingpin.Command("string", "Generates a random string of digits, upper and lower case characters.")
    stringLengthArg        = stringCommand.Arg("LENGTH", "Determines the length of the random data in bytes").Default("16").Int()
    uuidCommand            = kingpin.Command("uuid", "Generates a random UUID4.")
)

func main() {

    switch kingpin.Parse() {
    case "silly":

        // Print a random silly name
        fmt.Println(randomdata.SillyName())
    case "first":
        if *firstNameGender == "male" {
            // Print a male first name
            fmt.Println(randomdata.FirstName(randomdata.Male))

        } else if *firstNameGender == "female" {
            // Print a female first name
            fmt.Println(randomdata.FirstName(randomdata.Female))
        }
    case "last":

        // Print a last name
        fmt.Println(randomdata.LastName())
    case "full":
        if *fullNameGender == "male" {
            // Print a male first name
            fmt.Println(randomdata.FullName(randomdata.Male))

        } else if *fullNameGender == "female" {
            // Print a female first name
            fmt.Println(randomdata.FullName(randomdata.Female))
        }
    case "email":
        // Print an email
        fmt.Println(randomdata.Email())
    case "country":
        if *twoLetterCountryFlag {
            // Print a country using ISO 3166-1 alpha-2
            fmt.Println(randomdata.Country(randomdata.TwoCharCountry))
        } else if *threeLetterCountryFlag {
            // Print a country using ISO 3166-1 alpha-3
            fmt.Println(randomdata.Country(randomdata.ThreeCharCountry))
        } else {
            // Print a country with full text representation
            fmt.Println(randomdata.Country(randomdata.FullCountry))
        }
    case "currency":

        // Print a currency using ISO 4217
        fmt.Println(randomdata.Currency())
    case "city":

        // Print the name of a random city
        fmt.Println(randomdata.City())
    case "state":
        if *shortStateFlag {

            // Print the name of a random american state using two chars
            fmt.Println(randomdata.State(randomdata.Small))
        } else {

            // Print the name of a random american state
            fmt.Println(randomdata.State(randomdata.Large))
        }
    case "street":
        // Print an american sounding street name
        fmt.Println(randomdata.Street())
    case "address":
        // Print an american sounding address
        fmt.Println(randomdata.Address())
    case "number":
        // Print a random number >= 10 and <= 20
        fmt.Println(randomdata.Number(*startNumberArg, *stopNumberArg))
    case "decimal":
        // Print a random float >= 0 and <= 20 with decimal point 3
        fmt.Println(randomdata.Decimal(*startDecimalArg, *stopDecimalArg, *precDecimalArg))
    case "bool":
        // Print a bool
        fmt.Println(randomdata.Boolean())
    case "paragraph":
        // Print a paragraph
        fmt.Println(randomdata.Paragraph())
    case "postal":
        // Print a postal code
        fmt.Println(randomdata.PostalCode(*postalArg))
    case "ip":
        // Print a valid random IPv4 address
        fmt.Println(randomdata.IpV4Address())
    case "day":
        // Print a day
        fmt.Println(randomdata.Day())
    case "hex":
        var buf []byte
        if *hexLengthArg == 0 || *hexLengthArg > 16384 {
            buf = make([]byte, 16)
        } else {
            buf = make([]byte, *hexLengthArg)
        }
        rand.Read(buf)
        fmt.Println(hex.EncodeToString(buf))
    case "base64":
        var buf []byte
        if *base64LengthArg == 0 || *base64LengthArg > 16384 {
            buf = make([]byte, 16)
        } else {
            buf = make([]byte, *base64LengthArg)
        }
        rand.Read(buf)
        fmt.Println(base64.StdEncoding.EncodeToString(buf))
    case "string":
        var buf []byte
        if *stringLengthArg == 0 || *stringLengthArg > 16384 {
            buf = make([]byte, 16)
        } else {
            buf = make([]byte, *stringLengthArg)
        }
        r := RandStringReader{mrand.NewSource(time.Now().UnixNano())}
        r.Read(buf)
        fmt.Println(string(buf))
    case "uuid":
        fmt.Println(UUID4())
    }
}

type RandStringReader struct {
    src mrand.Source
}

func (r *RandStringReader) Read(p []byte) (int, error) {
    size := len(p)
    start := int64(r.src.Int63())
    for i := size - 1; i > -1; i-- {
        p[i] = StdChars[rune(start)&61]
        if i&7 == 0 {
            start = int64(r.src.Int63())
        }
        start >>= 7
    }
    return len(p), nil
}

var StdChars = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789")

// Zeros represents an empty or null UUID.
const Zeros = "00000000-0000-0000-0000-000000000000"

// UUID4 generates a random UUID according to RFC 4122
func UUID4() string {
    uuid := make([]byte, 16)
    n, err := io.ReadFull(rand.Reader, uuid)
    if n != len(uuid) || err != nil {
        return Zeros
    }

    // variant bits; see section 4.1.1
    uuid[8] = uuid[8]&^0xc0 | 0x80

    // version 4 (pseudo-random); see section 4.1.3
    uuid[6] = uuid[6]&^0xf0 | 0x40
    return fmt.Sprintf("%x-%x-%x-%x-%x", uuid[0:4], uuid[4:6], uuid[6:8], uuid[8:10], uuid[10:])
}
