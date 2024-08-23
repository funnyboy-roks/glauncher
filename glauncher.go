package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/bep/debounce"
	"github.com/mattn/go-gtk/gtk"
)

func apply_filter(labels []*gtk.Label, langs []string, filter string) {
    filled := 0;
    lower_filter := strings.ToLower(filter);
    for _, x := range langs {
        if filled >= len(labels) {
            break;
        }

        if strings.Contains(strings.ToLower(x), lower_filter) {
            labels[filled].SetText(x);
            filled += 1;
        }
    }

    for _, x := range labels[filled:] {
        x.SetText("");
    }
}

func main() {
    gtk.Init(nil)
    window := gtk.NewWindow(gtk.WINDOW_TOPLEVEL)
    window.SetPosition(gtk.WIN_POS_CENTER)
    window.SetTitle("GTK Go!")

    f, err := os.Open("./test-text.txt");

    if err != nil {
        log.Fatal("error:", err.Error());
    }

    defer f.Close();

    scan := bufio.NewScanner(f);

    langs := []string{};
    for scan.Scan() {
        langs = append(langs, scan.Text());
    }

    println("Language #:", len(langs));

    //--------------------------------------------------------
    // GtkVBox
    //--------------------------------------------------------
    vbox := gtk.NewVBox(false, 0)

    // label := gtk.NewLabel("Go Binding for GTK")
    // label.ModifyFontEasy("Anonymous Pro 24")
    // vbox.PackStart(label, false, true, 0)

    //--------------------------------------------------------
    // GtkEntry
    //--------------------------------------------------------
    entry := gtk.NewEntry()
    entry.SetText("Hello world")
    entry.ModifyFontEasy("monospace 24");
    vbox.PackStart(entry, false, true, 0)

    label := gtk.NewLabel("");
    vbox.PackStart(label, false, true, 0);

    const count = 20;
    labels := make([]*gtk.Label, count);

    changed_debounce := debounce.New(100 * time.Millisecond);
    entry.Connect("changed", func() {
        changed_debounce(func() {
            text := entry.GetBuffer().GetText();
            println("value:", text);
            // label.SetText(text);
            apply_filter(labels, langs, text);
        })
    });

    fmt.Println(langs);

    // frame := gtk.NewFrame("Langs")
    // framebox := gtk.NewVBox(false, 1)
    // frame.Add(framebox);
    // vbox.PackStart(frame, false, true, 0);
    for i, lang := range langs[:count] {
        labels[i] = gtk.NewLabel(lang)
        labels[i].ModifyFontEasy("Anonymous Pro 16")
        vbox.PackStart(labels[i], false, true, 0)
        println(lang);
    }

    // //--------------------------------------------------------
    // // GtkScale
    // //--------------------------------------------------------
    // scale := gtk.NewHScaleWithRange(0, 100, 1)
    // scale.Connect("value-changed", func() {
    //     println("scale:", int(scale.GetValue()))
    // })
    // vbox.Add(scale)

    // //--------------------------------------------------------
    // // GtkHBox
    // //--------------------------------------------------------
    // buttons := gtk.NewHBox(false, 1)

    // //--------------------------------------------------------
    // // GtkButton
    // //--------------------------------------------------------
    // button := gtk.NewButtonWithLabel("Button with label")
    // button.Clicked(func() {
    //     println("button clicked:", button.GetLabel())
    //     messagedialog := gtk.NewMessageDialog(
    //         button.GetTopLevelAsWindow(),
    //         gtk.DIALOG_MODAL,
    //         gtk.MESSAGE_INFO,
    //         gtk.BUTTONS_OK,
    //         entry.GetText())
    //     messagedialog.Response(func() {
    //         println("Dialog OK!")

    //         //--------------------------------------------------------
    //         // GtkFileChooserDialog
    //         //--------------------------------------------------------
    //         filechooserdialog := gtk.NewFileChooserDialog(
    //             "Choose File...",
    //             button.GetTopLevelAsWindow(),
    //             gtk.FILE_CHOOSER_ACTION_OPEN,
    //             gtk.STOCK_OK,
    //             gtk.RESPONSE_ACCEPT)
    //         filter := gtk.NewFileFilter()
    //         filter.AddPattern("*.go")
    //         filechooserdialog.AddFilter(filter)
    //         filechooserdialog.Response(func() {
    //             println(filechooserdialog.GetFilename())
    //             filechooserdialog.Destroy()
    //         })
    //         filechooserdialog.Run()
    //         messagedialog.Destroy()
    //     })
    //     messagedialog.Run()
    // })
    // buttons.Add(button)

    // //--------------------------------------------------------
    // // GtkFontButton
    // //--------------------------------------------------------
    // fontbutton := gtk.NewFontButton()
    // fontbutton.Connect("font-set", func() {
    //     println("title:", fontbutton.GetTitle())
    //     println("fontname:", fontbutton.GetFontName())
    //     println("use_size:", fontbutton.GetUseSize())
    //     println("show_size:", fontbutton.GetShowSize())
    // })
    // buttons.Add(fontbutton)
    // vbox.PackStart(buttons, false, false, 0)

    // buttons = gtk.NewHBox(false, 1)

    // //--------------------------------------------------------
    // // GtkToggleButton
    // //--------------------------------------------------------
    // togglebutton := gtk.NewToggleButtonWithLabel("ToggleButton with label")
    // togglebutton.Connect("toggled", func() {
    //     if togglebutton.GetActive() {
    //         togglebutton.SetLabel("ToggleButton ON!")
    //     } else {
    //         togglebutton.SetLabel("ToggleButton OFF!")
    //     }
    // })
    // buttons.Add(togglebutton)

    // //--------------------------------------------------------
    // // GtkCheckButton
    // //--------------------------------------------------------
    // checkbutton := gtk.NewCheckButtonWithLabel("CheckButton with label")
    // checkbutton.Connect("toggled", func() {
    //     if checkbutton.GetActive() {
    //         checkbutton.SetLabel("CheckButton CHECKED!")
    //     } else {
    //         checkbutton.SetLabel("CheckButton UNCHECKED!")
    //     }
    // })
    // buttons.Add(checkbutton)

    // //--------------------------------------------------------
    // // GtkRadioButton
    // //--------------------------------------------------------
    // buttonbox := gtk.NewVBox(false, 1)
    // radiofirst := gtk.NewRadioButtonWithLabel(nil, "Radio1")
    // buttonbox.Add(radiofirst)
    // buttonbox.Add(gtk.NewRadioButtonWithLabel(radiofirst.GetGroup(), "Radio2"))
    // buttonbox.Add(gtk.NewRadioButtonWithLabel(radiofirst.GetGroup(), "Radio3"))
    // buttons.Add(buttonbox)
    // //radiobutton.SetMode(false);
    // radiofirst.SetActive(true)

    // vbox.PackStart(buttons, false, false, 0)

    // //--------------------------------------------------------
    // // GtkVSeparator
    // //--------------------------------------------------------
    // vsep := gtk.NewVSeparator()
    // vbox.PackStart(vsep, false, false, 0)

    // //--------------------------------------------------------
    // // GtkComboBoxEntry
    // //--------------------------------------------------------
    // combos := gtk.NewHBox(false, 1)
    // comboboxentry := gtk.NewComboBoxEntryNewText()
    // comboboxentry.AppendText("Monkey")
    // comboboxentry.AppendText("Tiger")
    // comboboxentry.AppendText("Elephant")
    // comboboxentry.Connect("changed", func() {
    //     println("value:", comboboxentry.GetActiveText())
    // })
    // combos.Add(comboboxentry)

    // //--------------------------------------------------------
    // // GtkComboBox
    // //--------------------------------------------------------
    // combobox := gtk.NewComboBoxNewText()
    // combobox.AppendText("Peach")
    // combobox.AppendText("Banana")
    // combobox.AppendText("Apple")
    // combobox.SetActive(1)
    // combobox.Connect("changed", func() {
    //     println("value:", combobox.GetActiveText())
    // })
    // combos.Add(combobox)

    // vbox.PackStart(combos, false, false, 0)

    // //--------------------------------------------------------
    // // GtkTextView
    // //--------------------------------------------------------
    // swin := gtk.NewScrolledWindow(nil, nil)
    // swin.SetPolicy(gtk.POLICY_AUTOMATIC, gtk.POLICY_AUTOMATIC)
    // swin.SetShadowType(gtk.SHADOW_IN)
    // textview := gtk.NewTextView()
    // var start, end gtk.TextIter
    // buffer := textview.GetBuffer()
    // buffer.GetStartIter(&start)
    // buffer.Insert(&start, "Hello ")
    // buffer.GetEndIter(&end)
    // buffer.Insert(&end, "World!")
    // tag := buffer.CreateTag("bold", map[string]interface{}{
    //     "background": "#FF0000",
    //     "weight": "bold",
    // })
    // buffer.GetStartIter(&start)
    // buffer.GetEndIter(&end)
    // buffer.ApplyTag(tag, &start, &end)
    // swin.Add(textview)
    // vbox.Add(swin)

    // buffer.Connect("changed", func() {
    //     println("changed")
    // })

    window.Add(vbox)
    window.SetSizeRequest(600, 600)
    window.ShowAll()
    gtk.Main()
}
