// 由res2go自动生成，不要编辑。
package gui

import (
	"github.com/ying32/govcl/vcl"
)

type TMainForm struct {
	*vcl.TForm
	BtnFormat   *vcl.TButton
	Btn2        *vcl.TButton
	CBProtocols *vcl.TComboBox
	CB2         *vcl.TComboBox
	Panel1      *vcl.TPanel
	Splitter1   *vcl.TSplitter
	Panel2      *vcl.TPanel
	Memo1       *vcl.TMemo
	Panel3      *vcl.TPanel
	Memo2       *vcl.TMemo
	StringGrid1 *vcl.TStringGrid
	MainMenu    *vcl.TMainMenu
	ConnectItem *vcl.TMenuItem
	TestItem    *vcl.TMenuItem

	//::private::
	TMainFormFields
}

var MainForm *TMainForm

// 以字节形式加载
// vcl.Application.CreateForm(&MainForm)

func NewMainForm(owner vcl.IComponent) (root *TMainForm) {
	vcl.CreateResForm(owner, &root)
	return
}

var MainFormBytes = []byte("\x54\x50\x46\x30\x09\x54\x4D\x61\x69\x6E\x46\x6F\x72\x6D\x08\x4D\x61\x69\x6E\x46\x6F\x72\x6D\x04\x4C\x65\x66\x74\x03\x02\x02\x06\x48\x65\x69\x67\x68\x74\x03\x01\x02\x03\x54\x6F\x70\x02\x70\x05\x57\x69\x64\x74\x68\x03\xA8\x02\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x0B\x4D\x65\x73\x73\x61\x67\x65\x54\x6F\x6F\x6C\x0C\x43\x6C\x69\x65\x6E\x74\x48\x65\x69\x67\x68\x74\x03\xED\x01\x0B\x43\x6C\x69\x65\x6E\x74\x57\x69\x64\x74\x68\x03\xA8\x02\x04\x4D\x65\x6E\x75\x07\x08\x4D\x61\x69\x6E\x4D\x65\x6E\x75\x0A\x4C\x43\x4C\x56\x65\x72\x73\x69\x6F\x6E\x06\x08\x32\x2E\x30\x2E\x31\x30\x2E\x30\x00\x07\x54\x42\x75\x74\x74\x6F\x6E\x09\x42\x74\x6E\x46\x6F\x72\x6D\x61\x74\x04\x4C\x65\x66\x74\x03\xF8\x00\x06\x48\x65\x69\x67\x68\x74\x02\x19\x03\x54\x6F\x70\x02\x08\x05\x57\x69\x64\x74\x68\x02\x4B\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x06\x46\x6F\x72\x6D\x61\x74\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x00\x00\x00\x07\x54\x42\x75\x74\x74\x6F\x6E\x04\x42\x74\x6E\x32\x04\x4C\x65\x66\x74\x03\x4D\x01\x06\x48\x65\x69\x67\x68\x74\x02\x19\x03\x54\x6F\x70\x02\x08\x05\x57\x69\x64\x74\x68\x02\x4B\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x04\x42\x74\x6E\x32\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x01\x00\x00\x09\x54\x43\x6F\x6D\x62\x6F\x42\x6F\x78\x0B\x43\x42\x50\x72\x6F\x74\x6F\x63\x6F\x6C\x73\x04\x4C\x65\x66\x74\x02\x08\x06\x48\x65\x69\x67\x68\x74\x02\x19\x03\x54\x6F\x70\x02\x08\x05\x57\x69\x64\x74\x68\x02\x68\x0C\x41\x75\x74\x6F\x43\x6F\x6D\x70\x6C\x65\x74\x65\x09\x10\x41\x75\x74\x6F\x43\x6F\x6D\x70\x6C\x65\x74\x65\x54\x65\x78\x74\x0B\x0C\x63\x62\x61\x63\x74\x45\x6E\x61\x62\x6C\x65\x64\x16\x63\x62\x61\x63\x74\x45\x6E\x64\x4F\x66\x4C\x69\x6E\x65\x43\x6F\x6D\x70\x6C\x65\x74\x65\x14\x63\x62\x61\x63\x74\x53\x65\x61\x72\x63\x68\x41\x73\x63\x65\x6E\x64\x69\x6E\x67\x00\x0A\x49\x74\x65\x6D\x48\x65\x69\x67\x68\x74\x02\x11\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x02\x04\x54\x65\x78\x74\x06\x0E\x53\x65\x6C\x65\x63\x74\x50\x72\x6F\x74\x6F\x63\x6F\x6C\x00\x00\x09\x54\x43\x6F\x6D\x62\x6F\x42\x6F\x78\x03\x43\x42\x32\x04\x4C\x65\x66\x74\x02\x78\x06\x48\x65\x69\x67\x68\x74\x02\x19\x03\x54\x6F\x70\x02\x08\x05\x57\x69\x64\x74\x68\x02\x68\x0A\x49\x74\x65\x6D\x48\x65\x69\x67\x68\x74\x02\x11\x05\x53\x74\x79\x6C\x65\x07\x0E\x63\x73\x44\x72\x6F\x70\x44\x6F\x77\x6E\x4C\x69\x73\x74\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x03\x00\x00\x06\x54\x50\x61\x6E\x65\x6C\x06\x50\x61\x6E\x65\x6C\x31\x04\x4C\x65\x66\x74\x02\x08\x06\x48\x65\x69\x67\x68\x74\x03\x60\x01\x03\x54\x6F\x70\x02\x28\x05\x57\x69\x64\x74\x68\x03\x95\x02\x07\x41\x6E\x63\x68\x6F\x72\x73\x0B\x05\x61\x6B\x54\x6F\x70\x06\x61\x6B\x4C\x65\x66\x74\x07\x61\x6B\x52\x69\x67\x68\x74\x08\x61\x6B\x42\x6F\x74\x74\x6F\x6D\x00\x0A\x42\x65\x76\x65\x6C\x4F\x75\x74\x65\x72\x07\x06\x62\x76\x4E\x6F\x6E\x65\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x06\x50\x61\x6E\x65\x6C\x31\x0C\x43\x6C\x69\x65\x6E\x74\x48\x65\x69\x67\x68\x74\x03\x60\x01\x0B\x43\x6C\x69\x65\x6E\x74\x57\x69\x64\x74\x68\x03\x95\x02\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x04\x00\x09\x54\x53\x70\x6C\x69\x74\x74\x65\x72\x09\x53\x70\x6C\x69\x74\x74\x65\x72\x31\x04\x4C\x65\x66\x74\x03\xB4\x00\x06\x48\x65\x69\x67\x68\x74\x03\x60\x01\x03\x54\x6F\x70\x02\x00\x05\x57\x69\x64\x74\x68\x02\x05\x00\x00\x06\x54\x50\x61\x6E\x65\x6C\x06\x50\x61\x6E\x65\x6C\x32\x04\x4C\x65\x66\x74\x02\x00\x06\x48\x65\x69\x67\x68\x74\x03\x60\x01\x03\x54\x6F\x70\x02\x00\x05\x57\x69\x64\x74\x68\x03\xB4\x00\x05\x41\x6C\x69\x67\x6E\x07\x06\x61\x6C\x4C\x65\x66\x74\x0A\x42\x65\x76\x65\x6C\x4F\x75\x74\x65\x72\x07\x06\x62\x76\x4E\x6F\x6E\x65\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x06\x50\x61\x6E\x65\x6C\x32\x0C\x43\x6C\x69\x65\x6E\x74\x48\x65\x69\x67\x68\x74\x03\x60\x01\x0B\x43\x6C\x69\x65\x6E\x74\x57\x69\x64\x74\x68\x03\xB4\x00\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x01\x00\x05\x54\x4D\x65\x6D\x6F\x05\x4D\x65\x6D\x6F\x31\x04\x4C\x65\x66\x74\x02\x00\x06\x48\x65\x69\x67\x68\x74\x03\x60\x01\x03\x54\x6F\x70\x02\x00\x05\x57\x69\x64\x74\x68\x03\xB4\x00\x05\x41\x6C\x69\x67\x6E\x07\x08\x61\x6C\x43\x6C\x69\x65\x6E\x74\x0D\x4C\x69\x6E\x65\x73\x2E\x53\x74\x72\x69\x6E\x67\x73\x01\x06\x05\x4D\x65\x6D\x6F\x31\x00\x0A\x53\x63\x72\x6F\x6C\x6C\x42\x61\x72\x73\x07\x0A\x73\x73\x41\x75\x74\x6F\x42\x6F\x74\x68\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x00\x08\x57\x6F\x72\x64\x57\x72\x61\x70\x08\x00\x00\x00\x06\x54\x50\x61\x6E\x65\x6C\x06\x50\x61\x6E\x65\x6C\x33\x04\x4C\x65\x66\x74\x03\xB9\x00\x06\x48\x65\x69\x67\x68\x74\x03\x60\x01\x03\x54\x6F\x70\x02\x00\x05\x57\x69\x64\x74\x68\x03\xDC\x01\x05\x41\x6C\x69\x67\x6E\x07\x08\x61\x6C\x43\x6C\x69\x65\x6E\x74\x0A\x42\x65\x76\x65\x6C\x4F\x75\x74\x65\x72\x07\x06\x62\x76\x4E\x6F\x6E\x65\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x06\x50\x61\x6E\x65\x6C\x33\x0C\x43\x6C\x69\x65\x6E\x74\x48\x65\x69\x67\x68\x74\x03\x60\x01\x0B\x43\x6C\x69\x65\x6E\x74\x57\x69\x64\x74\x68\x03\xDC\x01\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x02\x00\x05\x54\x4D\x65\x6D\x6F\x05\x4D\x65\x6D\x6F\x32\x04\x4C\x65\x66\x74\x02\x00\x06\x48\x65\x69\x67\x68\x74\x03\x60\x01\x03\x54\x6F\x70\x02\x00\x05\x57\x69\x64\x74\x68\x03\xDC\x01\x05\x41\x6C\x69\x67\x6E\x07\x08\x61\x6C\x43\x6C\x69\x65\x6E\x74\x0D\x4C\x69\x6E\x65\x73\x2E\x53\x74\x72\x69\x6E\x67\x73\x01\x06\x05\x4D\x65\x6D\x6F\x32\x00\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x00\x00\x00\x00\x00\x0B\x54\x53\x74\x72\x69\x6E\x67\x47\x72\x69\x64\x0B\x53\x74\x72\x69\x6E\x67\x47\x72\x69\x64\x31\x04\x4C\x65\x66\x74\x02\x08\x06\x48\x65\x69\x67\x68\x74\x02\x56\x03\x54\x6F\x70\x03\x90\x01\x05\x57\x69\x64\x74\x68\x03\x95\x02\x08\x43\x6F\x6C\x43\x6F\x75\x6E\x74\x02\x03\x07\x43\x6F\x6C\x75\x6D\x6E\x73\x0E\x01\x0D\x54\x69\x74\x6C\x65\x2E\x43\x61\x70\x74\x69\x6F\x6E\x06\x05\x54\x69\x74\x6C\x65\x00\x01\x0D\x54\x69\x74\x6C\x65\x2E\x43\x61\x70\x74\x69\x6F\x6E\x06\x05\x54\x69\x74\x6C\x65\x05\x57\x69\x64\x74\x68\x03\x80\x00\x00\x01\x0D\x54\x69\x74\x6C\x65\x2E\x43\x61\x70\x74\x69\x6F\x6E\x06\x05\x54\x69\x74\x6C\x65\x05\x57\x69\x64\x74\x68\x03\xC2\x01\x00\x00\x0F\x44\x65\x66\x61\x75\x6C\x74\x43\x6F\x6C\x57\x69\x64\x74\x68\x02\x40\x09\x46\x69\x78\x65\x64\x43\x6F\x6C\x73\x02\x00\x09\x46\x69\x78\x65\x64\x52\x6F\x77\x73\x02\x00\x04\x46\x6C\x61\x74\x09\x07\x4F\x70\x74\x69\x6F\x6E\x73\x0B\x0F\x67\x6F\x46\x69\x78\x65\x64\x48\x6F\x72\x7A\x4C\x69\x6E\x65\x0A\x67\x6F\x48\x6F\x72\x7A\x4C\x69\x6E\x65\x0B\x67\x6F\x52\x6F\x77\x53\x65\x6C\x65\x63\x74\x0E\x67\x6F\x53\x6D\x6F\x6F\x74\x68\x53\x63\x72\x6F\x6C\x6C\x00\x0F\x52\x61\x6E\x67\x65\x53\x65\x6C\x65\x63\x74\x4D\x6F\x64\x65\x07\x08\x72\x73\x6D\x4D\x75\x6C\x74\x69\x08\x52\x6F\x77\x43\x6F\x75\x6E\x74\x02\x00\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x05\x0A\x54\x69\x74\x6C\x65\x53\x74\x79\x6C\x65\x07\x08\x74\x73\x4E\x61\x74\x69\x76\x65\x00\x00\x09\x54\x4D\x61\x69\x6E\x4D\x65\x6E\x75\x08\x4D\x61\x69\x6E\x4D\x65\x6E\x75\x00\x09\x54\x4D\x65\x6E\x75\x49\x74\x65\x6D\x0B\x43\x6F\x6E\x6E\x65\x63\x74\x49\x74\x65\x6D\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x07\x43\x6F\x6E\x6E\x65\x63\x74\x00\x00\x09\x54\x4D\x65\x6E\x75\x49\x74\x65\x6D\x08\x54\x65\x73\x74\x49\x74\x65\x6D\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x04\x54\x65\x73\x74\x00\x00\x00\x00")

// 注册Form资源
var _ = vcl.RegisterFormResource(MainForm, &MainFormBytes)
