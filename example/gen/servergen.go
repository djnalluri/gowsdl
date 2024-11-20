// Code generated by gowsdl DO NOT EDIT.

package gen

import (
	"encoding/xml"
	"errors"
	"fmt"
	"net/http"
	"reflect"
	"strings"
)

var wsdl = `<definitions name="StockQuote" targetNamespace="http://example.com/stockquote.wsdl" xmlns="http://schemas.xmlsoap.org/wsdl/" xmlns:soap="http://schemas.xmlsoap.org/wsdl/soap/" xmlns:tns="http://example.com/stockquote.wsdl" xmlns:xsd1="http://example.com/stockquote.xsd">
	<types>
		<schema targetNamespace="http://example.com/stockquote.xsd" xmlns="http://www.w3.org/2001/XMLSchema">
			<element name="TradePriceRequest">
				<complexType>
					<all>
						<element name="tickerSymbol" type="string"/>
					</all>
				</complexType>
			</element>
			<element name="TradePrice">
				<complexType>
					<all>
						<element name="price" type="float"/>
					</all>
				</complexType>
			</element>
		</schema>
	</types>
	<message name="GetLastTradePriceInput">
		<part element="xsd1:TradePriceRequest" name="body"/>
	</message>
	<message name="GetLastTradePriceOutput">
		<part element="xsd1:TradePrice" name="body"/>
	</message>
	<portType name="StockQuotePortType">
		<operation name="GetLastTradePrice">
			<input message="tns:GetLastTradePriceInput"/>
			<output message="tns:GetLastTradePriceOutput"/>
		</operation>
	</portType>
	<binding name="StockQuoteSoapBinding" type="tns:StockQuotePortType">
		<soap:binding style="document" transport="http://schemas.xmlsoap.org/soap/http"/>
		<operation name="GetLastTradePrice">
			<soap:operation soapAction="http://example.com/GetLastTradePrice"/>
			<input>
				<soap:body use="literal"/>
			</input>
			<output>
				<soap:body use="literal"/>
			</output>
		</operation>
	</binding>
	<service name="StockQuoteService">
		<documentation>My first service</documentation>
		<port binding="tns:StockQuoteBinding" name="StockQuotePort">
			<soap:address location="http://example.com/stockquote"/>
		</port>
	</service>
</definitions>
`

var WSDLUndefinedError = errors.New("Server was unable to process request. --> Object reference not set to an instance of an object.")

type SOAPEnvelopeRequest struct {
	XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Envelope"`
	Body    SOAPBodyRequest
}

type SOAPBodyRequest struct {
	XMLName xml.Name `xml:"http://schemas.xmlsoap.org/soap/envelope/ Body"`

	TradePriceRequest *TradePriceRequest `xml:,omitempty`
}

type SOAPEnvelopeResponse struct {
	XMLName    xml.Name `xml:"soap:Envelope"`
	PrefixSoap string   `xml:"xmlns:soap,attr"`
	PrefixXsi  string   `xml:"xmlns:xsi,attr"`
	PrefixXsd  string   `xml:"xmlns:xsd,attr"`

	Body SOAPBodyResponse
}

func NewSOAPEnvelopResponse() *SOAPEnvelopeResponse {
	return &SOAPEnvelopeResponse{
		PrefixSoap: "http://schemas.xmlsoap.org/soap/envelope/",
		PrefixXsd:  "http://www.w3.org/2001/XMLSchema",
		PrefixXsi:  "http://www.w3.org/2001/XMLSchema-instance",
	}
}

type Fault struct {
	XMLName xml.Name `xml:"SOAP-ENV:Fault"`
	Space   string   `xml:"xmlns:SOAP-ENV,omitempty,attr"`

	Code   string `xml:"faultcode,omitempty"`
	String string `xml:"faultstring,omitempty"`
	Actor  string `xml:"faultactor,omitempty"`
	Detail string `xml:"detail,omitempty"`
}

type SOAPBodyResponse struct {
	XMLName xml.Name `xml:"soap:Body"`
	Fault   *Fault   `xml:",omitempty"`

	TradePriceRequest *TradePrice `xml:",omitempty"`
}

func (service *SOAPBodyRequest) TradePriceRequestFunc(request *TradePriceRequest) (*TradePrice, error) {
	return nil, WSDLUndefinedError
}

func (service *SOAPEnvelopeRequest) call(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/xml; charset=utf-8")
	val := reflect.ValueOf(&service.Body).Elem()
	n := val.NumField()
	var field reflect.Value
	var name string
	find := false

	if r.Method == http.MethodGet {
		w.Write([]byte(wsdl))
		return
	}

	resp := NewSOAPEnvelopResponse()
	defer func() {
		if r := recover(); r != nil {
			resp.Body.Fault = &Fault{}
			resp.Body.Fault.Space = "http://schemas.xmlsoap.org/soap/envelope/"
			resp.Body.Fault.Code = "soap:Server"
			resp.Body.Fault.Detail = fmt.Sprintf("%v", r)
			resp.Body.Fault.String = fmt.Sprintf("%v", r)
		}
		xml.NewEncoder(w).Encode(resp)
	}()

	header := r.Header.Get("Content-Type")
	if strings.Index(header, "application/soap+xml") >= 0 {
		panic("Could not find an appropriate Transport Binding to invoke.")
	}

	err := xml.NewDecoder(r.Body).Decode(service)
	if err != nil {
		panic(err)
	}

	for i := 0; i < n; i++ {
		field = val.Field(i)
		name = val.Type().Field(i).Name
		if field.Kind() != reflect.Ptr {
			continue
		}
		if field.IsNil() {
			continue
		}
		if field.IsValid() {
			find = true
			break
		}
	}

	if !find {
		panic(WSDLUndefinedError)
	} else {
		m := val.Addr().MethodByName(name + "Func")
		if !m.IsValid() {
			panic(WSDLUndefinedError)
		}

		vals := m.Call([]reflect.Value{field})
		if vals[1].IsNil() {
			reflect.ValueOf(&resp.Body).Elem().FieldByName(name).Set(vals[0])
		} else {
			panic(vals[1].Interface())
		}
	}

}

func Endpoint(w http.ResponseWriter, r *http.Request) {
	request := SOAPEnvelopeRequest{}
	request.call(w, r)
}
