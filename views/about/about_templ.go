// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.680
package about

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func About() templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<style>\n        .img-container {\n            position: relative;\n            max-width: 20%; \n        }\n\n        .overlay-image {\n            display: block;\n            width: 100%;\n        }\n        .selected {\n          background-color: gray;\n        }\n\n        .overlay {\n          position: absolute;\n          bottom: 0;\n          left: 0;\n          right: 0;\n          background-color: black;\n          overflow: hidden;\n          width: 100%;\n          height: 0;\n          transition: .5s ease;\n        }\n\n        .img-container:hover .overlay {\n            bottom: 0; /* Show overlay on hover */\n            height: 100%;\n        }\n\n        .overlay-text {\n          color: white;\n          font-size: 20px;\n          position: absolute;\n          top: 50%;\n          left: 50%;\n          transform: translate(-50%, -50%);\n          text-align: center;\n        }\n    </style><div class=\"flex items-start justify-center h-full bg-black\"><div class=\"img-container flex items-center justify-center\"><img src=\"public/images/nereus_logo.png\" alt=\"Avatar\" class=\"overlay-image\"><div class=\"overlay\"><div class=\"overlay-text\"><div class=\"tab-list\" role=\"tablist\"><button hx-get=\"/about\" class=\"selected\" role=\"tab\" aria-selected=\"true\" aria-controls=\"tab-content\">About</button><hr><button hx-get=\"/blog\" role=\"tab\" aria-selected=\"false\" aria-controls=\"tab-content\">Blog</button><hr><button hx-get=\"/contact\" role=\"tab\" aria-selected=\"false\" aria-controls=\"tab-content\">Contact</button></div></div></div></div></div><div id=\"tab-content\" role=\"tabpanel\" class=\"tab-content flex items-center justify-center text-white\">Nereus, in Greek religion, sea god called by Homer “Old Man of the Sea,” noted for his wisdom, gift of prophecy, and ability to change his shape. He was the son of Pontus, a personification of the sea, and Gaea, the Earth goddess. The Nereids (water nymphs) were his daughters by the Oceanid Doris, and he lived with them in the depths of the sea, particularly the Aegean. Aphrodite, the goddess of love, was his pupil. The Greek hero Heracles, in his quest for the golden apples of the Hesperides, obtained directions from Nereus by wrestling with him in his many forms. Nereus frequently appears in vase paintings as a dignified spectator.</div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
