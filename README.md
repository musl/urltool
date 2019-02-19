# urltool
A simple tool to escape and unescape data for URLs.

# Build
`go build .`

# Use
`urltool escape < somefile > somefile_escaped`
`echo "Hello there." | urltool escape`
`echo "Hello%20there." | urltool unescape`

