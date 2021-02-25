package renderer

import(
	"testing"

	"github.com/alecthomas/chroma"
	"github.com/alecthomas/chroma/styles"
)

func TestRender(t *testing.T) {
	tests := []struct{
		md       string
		expected string
	}{
		{
			`# Header`,
			`<style>/* Background */ .chroma { color: #f8f8f2; background-color: #272822 }
/* Error */ .chroma .err { color: #960050; background-color: #1e0010 }
/* LineTableTD */ .chroma .lntd { vertical-align: top; padding: 0; margin: 0; border: 0; }
/* LineTable */ .chroma .lntable { border-spacing: 0; padding: 0; margin: 0; border: 0; width: auto; overflow: auto; display: block; }
/* LineHighlight */ .chroma .hl { display: block; width: 100%;background-color: #3c3d38 }
/* LineNumbersTable */ .chroma .lnt { margin-right: 0.4em; padding: 0 0.4em 0 0.4em;color: #7f7f7f }
/* LineNumbers */ .chroma .ln { margin-right: 0.4em; padding: 0 0.4em 0 0.4em;color: #7f7f7f }
/* Keyword */ .chroma .k { color: #66d9ef }
/* KeywordConstant */ .chroma .kc { color: #66d9ef }
/* KeywordDeclaration */ .chroma .kd { color: #66d9ef }
/* KeywordNamespace */ .chroma .kn { color: #f92672 }
/* KeywordPseudo */ .chroma .kp { color: #66d9ef }
/* KeywordReserved */ .chroma .kr { color: #66d9ef }
/* KeywordType */ .chroma .kt { color: #66d9ef }
/* NameAttribute */ .chroma .na { color: #a6e22e }
/* NameClass */ .chroma .nc { color: #a6e22e }
/* NameConstant */ .chroma .no { color: #66d9ef }
/* NameDecorator */ .chroma .nd { color: #a6e22e }
/* NameException */ .chroma .ne { color: #a6e22e }
/* NameFunction */ .chroma .nf { color: #a6e22e }
/* NameOther */ .chroma .nx { color: #a6e22e }
/* NameTag */ .chroma .nt { color: #f92672 }
/* Literal */ .chroma .l { color: #ae81ff }
/* LiteralDate */ .chroma .ld { color: #e6db74 }
/* LiteralString */ .chroma .s { color: #e6db74 }
/* LiteralStringAffix */ .chroma .sa { color: #e6db74 }
/* LiteralStringBacktick */ .chroma .sb { color: #e6db74 }
/* LiteralStringChar */ .chroma .sc { color: #e6db74 }
/* LiteralStringDelimiter */ .chroma .dl { color: #e6db74 }
/* LiteralStringDoc */ .chroma .sd { color: #e6db74 }
/* LiteralStringDouble */ .chroma .s2 { color: #e6db74 }
/* LiteralStringEscape */ .chroma .se { color: #ae81ff }
/* LiteralStringHeredoc */ .chroma .sh { color: #e6db74 }
/* LiteralStringInterpol */ .chroma .si { color: #e6db74 }
/* LiteralStringOther */ .chroma .sx { color: #e6db74 }
/* LiteralStringRegex */ .chroma .sr { color: #e6db74 }
/* LiteralStringSingle */ .chroma .s1 { color: #e6db74 }
/* LiteralStringSymbol */ .chroma .ss { color: #e6db74 }
/* LiteralNumber */ .chroma .m { color: #ae81ff }
/* LiteralNumberBin */ .chroma .mb { color: #ae81ff }
/* LiteralNumberFloat */ .chroma .mf { color: #ae81ff }
/* LiteralNumberHex */ .chroma .mh { color: #ae81ff }
/* LiteralNumberInteger */ .chroma .mi { color: #ae81ff }
/* LiteralNumberIntegerLong */ .chroma .il { color: #ae81ff }
/* LiteralNumberOct */ .chroma .mo { color: #ae81ff }
/* Operator */ .chroma .o { color: #f92672 }
/* OperatorWord */ .chroma .ow { color: #f92672 }
/* Comment */ .chroma .c { color: #75715e }
/* CommentHashbang */ .chroma .ch { color: #75715e }
/* CommentMultiline */ .chroma .cm { color: #75715e }
/* CommentSingle */ .chroma .c1 { color: #75715e }
/* CommentSpecial */ .chroma .cs { color: #75715e }
/* CommentPreproc */ .chroma .cp { color: #75715e }
/* CommentPreprocFile */ .chroma .cpf { color: #75715e }
/* GenericDeleted */ .chroma .gd { color: #f92672 }
/* GenericEmph */ .chroma .ge { font-style: italic }
/* GenericInserted */ .chroma .gi { color: #a6e22e }
/* GenericStrong */ .chroma .gs { font-weight: bold }
/* GenericSubheading */ .chroma .gu { color: #75715e }
</style><h1>Header</h1>
`,
		},
	}

	for _, tt := range tests {
		actual := Render(tt.md)
		if actual != tt.expected {
			t.Errorf("Render(%s) = %s, want %s", tt.md, actual, tt.expected)
		}
	}
}

func TestRenderWithStyle(t *testing.T) {
	tests := []struct{
		md       string
		style    *chroma.Style
		expected string
	}{
		{
			`# Header

` + "```vb" + `
Imports System

Public Module Hello
   Public Sub Main(  )
      Console.WriteLine("hello, world")
   End Sub
End Module
` + "```",
			styles.GitHub,
			`<style>/* Background */ .chroma { background-color: #ffffff }
/* Error */ .chroma .err { color: #a61717; background-color: #e3d2d2 }
/* LineTableTD */ .chroma .lntd { vertical-align: top; padding: 0; margin: 0; border: 0; }
/* LineTable */ .chroma .lntable { border-spacing: 0; padding: 0; margin: 0; border: 0; width: auto; overflow: auto; display: block; }
/* LineHighlight */ .chroma .hl { display: block; width: 100%;background-color: #e5e5e5 }
/* LineNumbersTable */ .chroma .lnt { margin-right: 0.4em; padding: 0 0.4em 0 0.4em;color: #7f7f7f }
/* LineNumbers */ .chroma .ln { margin-right: 0.4em; padding: 0 0.4em 0 0.4em;color: #7f7f7f }
/* Keyword */ .chroma .k { color: #000000; font-weight: bold }
/* KeywordConstant */ .chroma .kc { color: #000000; font-weight: bold }
/* KeywordDeclaration */ .chroma .kd { color: #000000; font-weight: bold }
/* KeywordNamespace */ .chroma .kn { color: #000000; font-weight: bold }
/* KeywordPseudo */ .chroma .kp { color: #000000; font-weight: bold }
/* KeywordReserved */ .chroma .kr { color: #000000; font-weight: bold }
/* KeywordType */ .chroma .kt { color: #445588; font-weight: bold }
/* NameAttribute */ .chroma .na { color: #008080 }
/* NameBuiltin */ .chroma .nb { color: #0086b3 }
/* NameBuiltinPseudo */ .chroma .bp { color: #999999 }
/* NameClass */ .chroma .nc { color: #445588; font-weight: bold }
/* NameConstant */ .chroma .no { color: #008080 }
/* NameDecorator */ .chroma .nd { color: #3c5d5d; font-weight: bold }
/* NameEntity */ .chroma .ni { color: #800080 }
/* NameException */ .chroma .ne { color: #990000; font-weight: bold }
/* NameFunction */ .chroma .nf { color: #990000; font-weight: bold }
/* NameLabel */ .chroma .nl { color: #990000; font-weight: bold }
/* NameNamespace */ .chroma .nn { color: #555555 }
/* NameTag */ .chroma .nt { color: #000080 }
/* NameVariable */ .chroma .nv { color: #008080 }
/* NameVariableClass */ .chroma .vc { color: #008080 }
/* NameVariableGlobal */ .chroma .vg { color: #008080 }
/* NameVariableInstance */ .chroma .vi { color: #008080 }
/* LiteralString */ .chroma .s { color: #dd1144 }
/* LiteralStringAffix */ .chroma .sa { color: #dd1144 }
/* LiteralStringBacktick */ .chroma .sb { color: #dd1144 }
/* LiteralStringChar */ .chroma .sc { color: #dd1144 }
/* LiteralStringDelimiter */ .chroma .dl { color: #dd1144 }
/* LiteralStringDoc */ .chroma .sd { color: #dd1144 }
/* LiteralStringDouble */ .chroma .s2 { color: #dd1144 }
/* LiteralStringEscape */ .chroma .se { color: #dd1144 }
/* LiteralStringHeredoc */ .chroma .sh { color: #dd1144 }
/* LiteralStringInterpol */ .chroma .si { color: #dd1144 }
/* LiteralStringOther */ .chroma .sx { color: #dd1144 }
/* LiteralStringRegex */ .chroma .sr { color: #009926 }
/* LiteralStringSingle */ .chroma .s1 { color: #dd1144 }
/* LiteralStringSymbol */ .chroma .ss { color: #990073 }
/* LiteralNumber */ .chroma .m { color: #009999 }
/* LiteralNumberBin */ .chroma .mb { color: #009999 }
/* LiteralNumberFloat */ .chroma .mf { color: #009999 }
/* LiteralNumberHex */ .chroma .mh { color: #009999 }
/* LiteralNumberInteger */ .chroma .mi { color: #009999 }
/* LiteralNumberIntegerLong */ .chroma .il { color: #009999 }
/* LiteralNumberOct */ .chroma .mo { color: #009999 }
/* Operator */ .chroma .o { color: #000000; font-weight: bold }
/* OperatorWord */ .chroma .ow { color: #000000; font-weight: bold }
/* Comment */ .chroma .c { color: #999988; font-style: italic }
/* CommentHashbang */ .chroma .ch { color: #999988; font-style: italic }
/* CommentMultiline */ .chroma .cm { color: #999988; font-style: italic }
/* CommentSingle */ .chroma .c1 { color: #999988; font-style: italic }
/* CommentSpecial */ .chroma .cs { color: #999999; font-weight: bold; font-style: italic }
/* CommentPreproc */ .chroma .cp { color: #999999; font-weight: bold; font-style: italic }
/* CommentPreprocFile */ .chroma .cpf { color: #999999; font-weight: bold; font-style: italic }
/* GenericDeleted */ .chroma .gd { color: #000000; background-color: #ffdddd }
/* GenericEmph */ .chroma .ge { color: #000000; font-style: italic }
/* GenericError */ .chroma .gr { color: #aa0000 }
/* GenericHeading */ .chroma .gh { color: #999999 }
/* GenericInserted */ .chroma .gi { color: #000000; background-color: #ddffdd }
/* GenericOutput */ .chroma .go { color: #888888 }
/* GenericPrompt */ .chroma .gp { color: #555555 }
/* GenericStrong */ .chroma .gs { font-weight: bold }
/* GenericSubheading */ .chroma .gu { color: #aaaaaa }
/* GenericTraceback */ .chroma .gt { color: #aa0000 }
/* GenericUnderline */ .chroma .gl { text-decoration: underline }
/* TextWhitespace */ .chroma .w { color: #bbbbbb }
</style><h1>Header</h1>
<pre style="background-color:#fff"><span style="color:#000;font-weight:bold">Imports</span> <span style="color:#555">System</span>

<span style="color:#000;font-weight:bold">Public</span> <span style="color:#000;font-weight:bold">Module</span> <span style="color:#555">Hello</span>
   <span style="color:#000;font-weight:bold">Public</span> <span style="color:#000;font-weight:bold">Sub</span> <span style="color:#900;font-weight:bold">Main</span>(  )
      Console.WriteLine(<span style="color:#d14">&#34;</span><span style="color:#d14">hello, world</span><span style="color:#d14">&#34;</span>)
   <span style="color:#000;font-weight:bold">End</span> <span style="color:#000;font-weight:bold">Sub</span>
<span style="color:#000;font-weight:bold">End</span> <span style="color:#000;font-weight:bold">Module</span>
</pre>`,
		},
	}

	for _, tt := range tests {
		actual := RenderWithStyle(tt.md, tt.style)
		if actual != tt.expected {
			t.Errorf("Render(%s) = %s, want %s", tt.md, actual, tt.expected)
		}
	}
}
