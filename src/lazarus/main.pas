unit main;

{$mode objfpc}{$H+}

interface

uses
  Classes, SysUtils, Forms, Controls, Graphics, Dialogs, ActnList, StdCtrls,
  Menus, ComCtrls, Buttons, ExtCtrls;

type

  { TMainForm }

  TMainForm = class(TForm)
    Button1: TButton;
    MainMenu: TMainMenu;
    Memo1: TMemo;
    ConnectItem: TMenuItem;
    TestItem: TMenuItem;
  private

  public

  end;

var
  MainForm: TMainForm;

implementation

{$R *.lfm}

end.

