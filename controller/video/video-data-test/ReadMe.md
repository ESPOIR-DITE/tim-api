#### Video-Data-Controller

This is a very critical controller and a bit complicated because there's enough mechanics.

First the VideoData Object
Contains Video field that is omitted in the Database. This is because, video are not supposed to be saved in the database. They may be saved in either in S3 or other external storage.

`
type VideoData struct { Id string Picture []byte Video []byte FileType string FileSize string }
`
