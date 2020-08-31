// 由res2go自动生成，不要编辑。
package gui

import (
	"github.com/ying32/govcl/vcl"
)

type TMainForm struct {
	*vcl.TForm
	Memo1       *vcl.TMemo
	Button1     *vcl.TButton
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

var MainFormBytes = []byte("\x54\x50\x46\x30\x09\x54\x4D\x61\x69\x6E\x46\x6F\x72\x6D\x08\x4D\x61\x69\x6E\x46\x6F\x72\x6D\x04\x4C\x65\x66\x74\x03\x9D\x02\x06\x48\x65\x69\x67\x68\x74\x03\x4F\x01\x03\x54\x6F\x70\x02\x72\x05\x57\x69\x64\x74\x68\x03\xA1\x01\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x0B\x4D\x65\x73\x73\x61\x67\x65\x54\x6F\x6F\x6C\x0C\x43\x6C\x69\x65\x6E\x74\x48\x65\x69\x67\x68\x74\x03\x3B\x01\x0B\x43\x6C\x69\x65\x6E\x74\x57\x69\x64\x74\x68\x03\xA1\x01\x04\x4D\x65\x6E\x75\x07\x08\x4D\x61\x69\x6E\x4D\x65\x6E\x75\x0A\x4C\x43\x4C\x56\x65\x72\x73\x69\x6F\x6E\x06\x08\x32\x2E\x30\x2E\x31\x30\x2E\x30\x00\x05\x54\x4D\x65\x6D\x6F\x05\x4D\x65\x6D\x6F\x31\x04\x4C\x65\x66\x74\x02\x08\x06\x48\x65\x69\x67\x68\x74\x03\x08\x01\x03\x54\x6F\x70\x02\x28\x05\x57\x69\x64\x74\x68\x03\x90\x01\x0D\x4C\x69\x6E\x65\x73\x2E\x53\x74\x72\x69\x6E\x67\x73\x01\x06\x05\x4D\x65\x6D\x6F\x31\x00\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x00\x00\x00\x07\x54\x42\x75\x74\x74\x6F\x6E\x07\x42\x75\x74\x74\x6F\x6E\x31\x04\x4C\x65\x66\x74\x02\x08\x06\x48\x65\x69\x67\x68\x74\x02\x19\x03\x54\x6F\x70\x02\x08\x05\x57\x69\x64\x74\x68\x02\x4B\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x07\x42\x75\x74\x74\x6F\x6E\x31\x08\x54\x61\x62\x4F\x72\x64\x65\x72\x02\x01\x00\x00\x09\x54\x4D\x61\x69\x6E\x4D\x65\x6E\x75\x08\x4D\x61\x69\x6E\x4D\x65\x6E\x75\x00\x09\x54\x4D\x65\x6E\x75\x49\x74\x65\x6D\x0B\x43\x6F\x6E\x6E\x65\x63\x74\x49\x74\x65\x6D\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x07\x43\x6F\x6E\x6E\x65\x63\x74\x00\x00\x09\x54\x4D\x65\x6E\x75\x49\x74\x65\x6D\x08\x54\x65\x73\x74\x49\x74\x65\x6D\x07\x43\x61\x70\x74\x69\x6F\x6E\x06\x04\x54\x65\x73\x74\x00\x00\x00\x00")

// 注册Form资源
var _ = vcl.RegisterFormResource(MainForm, &MainFormBytes)
