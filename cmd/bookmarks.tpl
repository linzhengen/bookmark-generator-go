<!DOCTYPE NETSCAPE-Bookmark-file-1>
<META HTTP-EQUIV="Content-Type" CONTENT="text/html; charset=UTF-8">
<TITLE>Bookmarks</TITLE>
<H1>Bookmarks</H1>
<DL><p>
    <DT><H3 PERSONAL_TOOLBAR_FOLDER="true">Bookmark Bar</H3>
    <DL><p>
      {{ range .Folders }}
        <DT><H3 LAST_MODIFIED="1621598878">{{ .Name }}</H3>
        {{ range .Bookmarks }}
        <DL><p>
            <DT><A HREF="{{ .URL }}" ADD_DATE="{{ .AddDate }}" ICON="{{ .IconURL }}">{{ .Name }}</A>
        </DL><p>
        {{ end }}
      {{ end }}
      {{ range .Bookmarks }}
        <DT><A HREF="{{ .URL }}" ADD_DATE="{{ .AddDate }}" ICON="{{ .IconURL }}">{{ .Name }}</A>
      {{ end }}
    </DL><p>
</DL><p>