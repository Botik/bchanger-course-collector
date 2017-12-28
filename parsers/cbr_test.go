package parsers

import "testing"

var cbrData = []byte(`<?xml version="1.0" encoding="windows-1251" ?>
	<ValCurs Date="20.12.2017" name="Foreign Currency Market">
	<Valute ID="R01010">
		<NumCode>036</NumCode>
		<CharCode>AUD</CharCode>
		<Nominal>1</Nominal>
		<Name>������������� ������</Name>
		<Value>44,9962</Value>
	</Valute>
	<Valute ID="R01020A">
		<NumCode>944</NumCode>
		<CharCode>AZN</CharCode>
		<Nominal>1</Nominal>
		<Name>��������������� �����</Name>
		<Value>34,4167</Value>
	</Valute>
	<Valute ID="R01035">
		<NumCode>826</NumCode>
		<CharCode>GBP</CharCode>
		<Nominal>1</Nominal>
		<Name>���� ���������� ������������ �����������</Name>
		<Value>78,4752</Value>
	</Valute>
	<Valute ID="R01060">
		<NumCode>051</NumCode>
		<CharCode>AMD</CharCode>
		<Nominal>100</Nominal>
		<Name>��������� ������</Name>
		<Value>12,1349</Value>
	</Valute>
	<Valute ID="R01090B">
		<NumCode>933</NumCode>
		<CharCode>BYN</CharCode>
		<Nominal>1</Nominal>
		<Name>����������� �����</Name>
		<Value>29,0877</Value>
	</Valute>
	<Valute ID="R01100">
		<NumCode>975</NumCode>
		<CharCode>BGN</CharCode>
		<Nominal>1</Nominal>
		<Name>���������� ���</Name>
		<Value>35,3935</Value>
	</Valute>
	<Valute ID="R01115">
		<NumCode>986</NumCode>
		<CharCode>BRL</CharCode>
		<Nominal>1</Nominal>
		<Name>����������� ����</Name>
		<Value>17,8027</Value>
	</Valute>
	<Valute ID="R01135">
		<NumCode>348</NumCode>
		<CharCode>HUF</CharCode>
		<Nominal>100</Nominal>
		<Name>���������� ��������</Name>
		<Value>22,1226</Value>
	</Valute>
	<Valute ID="R01200">
		<NumCode>344</NumCode>
		<CharCode>HKD</CharCode>
		<Nominal>10</Nominal>
		<Name>����������� ��������</Name>
		<Value>74,9481</Value>
	</Valute>
	<Valute ID="R01215">
		<NumCode>208</NumCode>
		<CharCode>DKK</CharCode>
		<Nominal>10</Nominal>
		<Name>������� ����</Name>
		<Value>93,0093</Value>
	</Valute>
	<Valute ID="R01235">
		<NumCode>840</NumCode>
		<CharCode>USD</CharCode>
		<Nominal>1</Nominal>
		<Name>������ ���</Name>
		<Value>58,6117</Value>
	</Valute>
	<Valute ID="R01239">
		<NumCode>978</NumCode>
		<CharCode>EUR</CharCode>
		<Nominal>1</Nominal>
		<Name>����</Name>
		<Value>69,1735</Value>
	</Valute>
	<Valute ID="R01270">
		<NumCode>356</NumCode>
		<CharCode>INR</CharCode>
		<Nominal>100</Nominal>
		<Name>��������� �����</Name>
		<Value>91,5164</Value>
	</Valute>
	<Valute ID="R01335">
		<NumCode>398</NumCode>
		<CharCode>KZT</CharCode>
		<Nominal>100</Nominal>
		<Name>������������� �����</Name>
		<Value>17,4624</Value>
	</Valute>
	<Valute ID="R01350">
		<NumCode>124</NumCode>
		<CharCode>CAD</CharCode>
		<Nominal>1</Nominal>
		<Name>��������� ������</Name>
		<Value>45,5803</Value>
	</Valute>
	<Valute ID="R01370">
		<NumCode>417</NumCode>
		<CharCode>KGS</CharCode>
		<Nominal>100</Nominal>
		<Name>���������� �����</Name>
		<Value>84,2122</Value>
	</Valute>
	<Valute ID="R01375">
		<NumCode>156</NumCode>
		<CharCode>CNY</CharCode>
		<Nominal>10</Nominal>
		<Name>��������� �����</Name>
		<Value>88,7061</Value>
	</Valute>
	<Valute ID="R01500">
		<NumCode>498</NumCode>
		<CharCode>MDL</CharCode>
		<Nominal>10</Nominal>
		<Name>���������� ����</Name>
		<Value>34,2258</Value>
	</Valute>
	<Valute ID="R01535">
		<NumCode>578</NumCode>
		<CharCode>NOK</CharCode>
		<Nominal>10</Nominal>
		<Name>���������� ����</Name>
		<Value>70,1080</Value>
	</Valute>
	<Valute ID="R01565">
		<NumCode>985</NumCode>
		<CharCode>PLN</CharCode>
		<Nominal>1</Nominal>
		<Name>�������� ������</Name>
		<Value>16,4829</Value>
	</Valute>
	<Valute ID="R01585F">
		<NumCode>946</NumCode>
		<CharCode>RON</CharCode>
		<Nominal>1</Nominal>
		<Name>��������� ���</Name>
		<Value>14,9956</Value>
	</Valute>
	<Valute ID="R01589">
		<NumCode>960</NumCode>
		<CharCode>XDR</CharCode>
		<Nominal>1</Nominal>
		<Name>��� (����������� ����� �������������)</Name>
		<Value>82,8207</Value>
	</Valute>
	<Valute ID="R01625">
		<NumCode>702</NumCode>
		<CharCode>SGD</CharCode>
		<Nominal>1</Nominal>
		<Name>������������ ������</Name>
		<Value>43,5225</Value>
	</Valute>
	<Valute ID="R01670">
		<NumCode>972</NumCode>
		<CharCode>TJS</CharCode>
		<Nominal>10</Nominal>
		<Name>���������� ������</Name>
		<Value>66,4909</Value>
	</Valute>
	<Valute ID="R01700J">
		<NumCode>949</NumCode>
		<CharCode>TRY</CharCode>
		<Nominal>1</Nominal>
		<Name>�������� ����</Name>
		<Value>15,2933</Value>
	</Valute>
	<Valute ID="R01710A">
		<NumCode>934</NumCode>
		<CharCode>TMT</CharCode>
		<Nominal>1</Nominal>
		<Name>����� ����������� �����</Name>
		<Value>16,7702</Value>
	</Valute>
	<Valute ID="R01717">
		<NumCode>860</NumCode>
		<CharCode>UZS</CharCode>
		<Nominal>10000</Nominal>
		<Name>��������� �����</Name>
		<Value>72,3458</Value>
	</Valute>
	<Valute ID="R01720">
		<NumCode>980</NumCode>
		<CharCode>UAH</CharCode>
		<Nominal>10</Nominal>
		<Name>���������� ������</Name>
		<Value>21,0228</Value>
	</Valute>
	<Valute ID="R01760">
		<NumCode>203</NumCode>
		<CharCode>CZK</CharCode>
		<Nominal>10</Nominal>
		<Name>������� ����</Name>
		<Value>26,9343</Value>
	</Valute>
	<Valute ID="R01770">
		<NumCode>752</NumCode>
		<CharCode>SEK</CharCode>
		<Nominal>10</Nominal>
		<Name>�������� ����</Name>
		<Value>69,6498</Value>
	</Valute>
	<Valute ID="R01775">
		<NumCode>756</NumCode>
		<CharCode>CHF</CharCode>
		<Nominal>1</Nominal>
		<Name>����������� �����</Name>
		<Value>59,4982</Value>
	</Valute>
	<Valute ID="R01810">
		<NumCode>710</NumCode>
		<CharCode>ZAR</CharCode>
		<Nominal>10</Nominal>
		<Name>��������������� ������</Name>
		<Value>46,0006</Value>
	</Valute>
	<Valute ID="R01815">
		<NumCode>410</NumCode>
		<CharCode>KRW</CharCode>
		<Nominal>1000</Nominal>
		<Name>��� ���������� �����</Name>
		<Value>54,1238</Value>
	</Valute>
	<Valute ID="R01820">
		<NumCode>392</NumCode>
		<CharCode>JPY</CharCode>
		<Nominal>100</Nominal>
		<Name>�������� ���</Name>
		<Value>52,0553</Value>
	</Valute>
	</ValCurs>`)

func BenchmarkParseCbrData(b *testing.B) {
	b.SetBytes(int64(len(cbrData)))

	for i := 0; i < b.N; i++ {
		if _, err := ParseCbrData(cbrData); nil != err {
			b.Error(err)
		}
	}
}
