package gomailify

// HTMLTemplate taken from react-email package.
// Combines both HTML and Head template
const HTMLTemplate = `
<!DOCTYPE html PUBLIC "-//W3C//DTD XHTML 1.0 Transitional//EN" "http://www.w3.org/TR/xhtml1/DTD/xhtml1-transitional.dtd">
<html dir="ltr" lang="en">
  <head>
    <meta content="text/html; charset=UTF-8" http-equiv="Content-Type" />
    <meta name="x-apple-disable-message-reformatting" />
</head>
    {{ .Value }}
    {{ range .Children }}
        {{ . }}
    {{ end }}
</html>
`

// ContainerTemplate taken from rendered react-email package.
const ContainerTemplate = `
<table
  align="center"
  width="100%"
  border="0"
  cellpadding="0"
  cellspacing="0"
  role="presentation"
  style="max-width:37.5em"
>
  <tbody>
    <tr style="width: 100%;">
      <td>
        {{ .Value }}
        {{ range .Children }}
          {{ . }}
        {{ end }}
      </td>
    </tr>
  </tbody>
</table>
`

// ParagraphTemplate taken from rendered react-email package.
const ParagraphTemplate = `
<p style="font-size:14px;line-height:24px;margin:16px 0">
    {{ .Value }}
    {{ range .Children }}
        {{ . }}
    {{ end }}
</p>
`

const CodeBlockTemplate = `
<table
  align="center"
  width="100%"
  border="0"
  cellpadding="0"
  cellspacing="0"
  role="presentation"
  style="max-width:37.5em"
>
  <tbody>
    <tr style="width: 100%;">
      <td style="font-family:monospace,'Courier New',Consolas,Courier,Monaco,'Andale Mono','Ubuntu Mono';">
		{{ .Value }}
		{{ range .Children }}
			{{ . }}
		{{ end }}
      </td>
    </tr>
  </tbody>
</table>
`
