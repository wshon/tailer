// 由res2go自动生成，不要编辑。
package gui

import (
	"github.com/ying32/govcl/vcl"
)

type TConnectForm struct {
	*vcl.TForm
	ListBox1  *vcl.TListBox
	Panel1    *vcl.TPanel
	ComboBox1 *vcl.TComboBox

	//::private::
	TConnectFormFields
}

var ConnectForm *TConnectForm

// 以字节形式加载
// vcl.Application.CreateForm(&ConnectForm)

func NewConnectForm(owner vcl.IComponent) (root *TConnectForm) {
	vcl.CreateResForm(owner, &root)
	return
}

var ConnectFormBytes = []byte("\x54\x50\x46\x30\x0C\x54\x43\x6F\x6E\x6E\x65\x63\x74\x46\x6F\x72\x6D\x0B\x43\x6F\x6E\x6E\x65\x63\x74\x46\x6F\x72\x6D\x04\x4C\x65\x66\x74\x03\x4F\x04\x06\x48\x65\x69\x67\x68\x74\x03\xF0\x00\x03\x54\x6F\x70\x02\x73\x05\x57\x69\x64\x74\x68\x03\x40\x01\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x0D\x43\x6F\x6E\x6E\x65\x63\x74\x4D\x61\x6E\x61\x67\x65\x0C\x43\x6C\x69\x65\x6E\x74\x48\x65\x69\x67\x68\x74\x03\xF0\x00\x0B\x43\x6C\x69\x65\x6E\x74\x57\x69\x64\x74\x68\x03\x40\x01\x0A\x4C\x43\x4C\x56\x65\x72\x73\x69\x6F\x6E\x06\x08\x32\x2E\x30\x2E\x31\x30\x2E\x30\x00\x08\x54\x4C\x69\x73\x74\x42\x6F\x78\x08\x4C\x69\x73\x74\x42\x6F\x78\x31\x04\x4C\x65\x66\x74\x02\x08\x06\x48\x65\x69\x67\x68\x74\x03\xE0\x00\x03\x54\x6F\x70\x02\x08\x05\x57\x69\x64\x74\x68\x02\x64\x0A\x49\x74\x65\x6D\x48\x65\x69\x67\x68\x74\x02\x00\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x00\x00\x00\x06\x54\x50\x61\x6E\x65\x6C\x06\x50\x61\x6E\x65\x6C\x31\x04\x4C\x65\x66\x74\x02\x70\x06\x48\x65\x69\x67\x68\x74\x03\xE2\x00\x03\x54\x6F\x70\x02\x08\x05\x57\x69\x64\x74\x68\x03\xCA\x00\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x06\x50\x61\x6E\x65\x6C\x31\x0C\x43\x6C\x69\x65\x6E\x74\x48\x65\x69\x67\x68\x74\x03\xE2\x00\x0B\x43\x6C\x69\x65\x6E\x74\x57\x69\x64\x74\x68\x03\xCA\x00\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x01\x00\x09\x54\x43\x6F\x6D\x62\x6F\x42\x6F\x78\x09\x43\x6F\x6D\x62\x6F\x42\x6F\x78\x31\x04\x4C\x65\x66\x74\x02\x60\x06\x48\x65\x69\x67\x68\x74\x02\x19\x03\x54\x6F\x70\x02\x08\x05\x57\x69\x64\x74\x68\x02\x64\x0A\x49\x74\x65\x6D\x48\x65\x69\x67\x68\x74\x02\x11\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x00\x04\x54\x65\x78\x74\x06\x09\x43\x6F\x6D\x62\x6F\x42\x6F\x78\x31\x00\x00\x00\x00")

// 注册Form资源
var _ = vcl.RegisterFormResource(ConnectForm, &ConnectFormBytes)
