# FFMPEG

https://medium.com/learning-the-go-programming-language/realtime-video-capture-with-go-65a8ac3a57da
https://github.com/amenzhinsky/go-memexec

https://github.com/pion/mediadevices

## list devices,
ffmpeg -f avfoundation -list_devices true -i ""

## record webcam
ffmpeg -f avfoundation -pix_fmt uyvy422 -framerate 30 -i "0:0" -y -target film-dvd -s 1280x720 -r 10 ./test2.mpg -hide_banner
ffmpeg -f avfoundation -framerate 30  -i "default" webcam.mpg

ffmpeg -f avfoundation -i "0" webcam.mpg

## record screen
ffmpeg -f avfoundation -pix_fmt uyvy422 -i "2" -t 20 -r 10 -codec copy -map 0 -f segment -segment_time 4 screen_%03d.mkv

ffmpeg -f avfoundation -pix_fmt uyvy422 -i "2" -t 5 -r 10 -f framehash out.sha256

## HUD
ffmpeg -f avfoundation -i "2"
ffmpeg -f avfoundation -framerate 30 -i "default" -vf scale=380:-1 out.mov
ffmpeg -f avfoundation -i "2" -f avfoundation -framerate 30 -i "default" -filter_complex "[1:v]scale=640:-1[scaled];[0:v][scaled]overlay=W-w-20:H-h-20[out]" -map "[out]" out.mkv
ffmpeg -f avfoundation -i "2" -f avfoundation -framerate 30 -i "default" -filter_complex "[1:v]scale=640:-1[scaled];[0:v][scaled]overlay=W-w-20:H-h-20[out]" -map "[out]" -c:v h264_videotoolbox -realtime 1 -profile:v 3 -level 51 -threads 4 -b:v 4000k output.mkv
ffmpeg -f avfoundation -i "2" -f avfoundation -framerate 30 -i "default" -filter_complex "[1:v]scale=640:-1[scaled];[0:v][scaled]overlay=W-w-20:H-h-20[out]" -map "[out]" -c:v h264_videotoolbox -realtime 1 -profile:v 3 -level 51 -threads 4 -b:v 4000k -r 10 -map 0 -f segment -segment_time 05:00 screen_%03d.mkv
ffmpeg -f avfoundation -i "2" -f avfoundation -framerate 30 -i "default" -filter_complex "[1:v]scale=640:-1[scaled];[0:v][scaled]overlay=W-w-20:H-h-20,split=2[overlay0][overlay1];[overlay0]scale=1920:-2[out0];[overlay1]scale=640:-2[out1]" -map "[out0]" -c:v h264_videotoolbox -realtime 1 -profile:v 3 -level 51 -threads 4 -b:v 4000k -r 10 -map 0 -f segment -segment_time 05:00 screen_%03d.mkv -map "[out1]" -b:v 2000k -r 10 -f mpegts udp://127.0.0.1:1234
ffmpeg -f avfoundation -i "2" -f avfoundation -framerate 30 -i "default" -filter_complex "[1:v]scale=640:-1[scaled];[0:v][scaled]overlay=W-w-20:H-h-20,split=2[overlay0][overlay1];[overlay0]scale=1920:-2[out0];[overlay1]scale=1024 :-2[out1]" -map "[out0]" -c:v h264_videotoolbox -realtime 1 -profile:v 3 -level 51 -threads 4 -b:v 4000k -r 10 -map 0 -f segment -segment_time 05:00 screen_%03d.mkv -map "[out1]" -b:v 2000k -r 10 -f mpegts udp://127.0.0.1:1234
ffmpeg -f avfoundation -i "2" -f avfoundation -framerate 30 -i "default" -filter_complex "[1:v]scale=640:-1[scaled];[0:v][scaled]overlay=W-w-20:H-h-20,drawtext=fontfile=arial.ttf: fontsize=24: fontcolor=yellow: text='%{localtime\:%a %d %b %Y %T} - %{pts\:hms}': box=1: boxcolor=black@0.5: boxborderw=5: x=7: y=7,split=2[overlay0][overlay1];[overlay0]scale=1920:-2[out0];[overlay1]scale=640 :-2[out1]" -map "[out0]" -c:v h264_videotoolbox -realtime 1 -profile:v 3 -level 51 -threads 4 -b:v 4000k -r 10 -f segment -segment_time 05:00 screen_%03d.mkv -map "[out1]" -b:v 2000k -r 24 -f mpegts udp://127.0.0.1:1234

### original size
ffmpeg -f avfoundation -capture_cursor 1 -capture_mouse_clicks 1 -framerate 5 -i "3" -f avfoundation -framerate 30 -video_size 640x480 -i "default" -filter_complex "[0:v][1:v]overlay=W-w-20:H-h-20,fps=fps=10,drawtext=fontfile=arial.ttf: fontsize=24: fontcolor=yellow: text='%{localtime\:%a %d %b %Y %T} - %{pts\:hms}': box=1: boxcolor=black@0.5: boxborderw=5: x=10: y=10[out]" -map "[out]" -c:v h264_videotoolbox -realtime 1 -profile:v 3 -level 51 -b:v 4000k -f segment -segment_time 05:00 -reset_timestamps 1 -segment_list videos/screen_list.csv videos/screen_%03d.mkv

### scaled
ffmpeg -f avfoundation -capture_cursor 1 -capture_mouse_clicks 1 -framerate 10 -i "3" -f avfoundation -framerate 30 -video_size 640x480 -i "default" -filter_complex "[0:v][1:v]overlay=W-w-20:H-h-20,fps=fps=10,scale=1280:-1,drawtext=fontfile=arial.ttf: fontsize=24: fontcolor=yellow: text='%{localtime\:%a %d %b %Y %T} - %{pts\:hms} - %{frame_num}': box=1: boxcolor=black@0.5: boxborderw=5: x=10: y=10[out]" -map "[out]" -c:v h264_videotoolbox -realtime 1 -profile:v 3 -level 51 -b:v 4000k -f segment -segment_time 05:00 -reset_timestamps 1 -strftime 1 videos/screen_%d%m%y_%H%M.mkv
ffmpeg -t 20 -f avfoundation -capture_cursor 1 -capture_mouse_clicks 1 -framerate 10 -i "3" -f avfoundation -framerate 30 -video_size 640x480 -i "default" -filter_complex "[0:v][1:v]overlay=W-w-20:H-h-20,fps=fps=10,scale=1280:-1,drawtext=fontfile=arial.ttf: fontsize=24: fontcolor=yellow: text='%{localtime\:%a %d %b %Y %T} - %{pts\:hms} - %{frame_num}': box=1: boxcolor=black@0.5: boxborderw=5: x=10: y=10[out]" -map "[out]" -c:v h264_videotoolbox -realtime 1 -profile:v 3 -level 51 -b:v 4000k videos/screen_0.mkv

ffmpeg -f avfoundation -capture_cursor 1 -capture_mouse_clicks 1 -framerate 5 -formats  -i "2"

steganography

## PROBE

ffprobe -show_format screen_000.mkv
ffprobe -show_streams -count_packets -count_frames screen_000.mkv
ffprobe -show_packets screen_000.mkv
ffprobe -show_packets screen_000.mkv
ffprobe -show_frames screen_000.mkv
ffprobe -show_packets -show_data_hash SHA256 -print_format xml screen_000.mkv
ffprobe -show_entries "packet=pts,duration,pos,size,data_hash" -bitexact -show_data_hash SHA256 -print_format xml screen_000.mkv
ffprobe -show_entries "frame=coded_picture_number,key_frame,pts,pkt_duration,pkt_pos,pkt_size" -bitexact -show_data_hash SHA256 -print_format xml screen_000.mkv

## virtual filesystem

https://blog.gopheracademy.com/advent-2014/fuse-zipfs/
https://askubuntu.com/questions/424201/how-can-i-protect-a-file-from-user-changes

chflags uchg out.mkv
sudo chflags schg fileName

## multiple diffusion
ffmpeg -f avfoundation -i "2" -filter_complex "[0:v]split=2[overlay0][overlay1];[overlay0]scale=1920:-2[out0];[overlay1]scale=640:-2[out1]" -map "[out0]" -c:v h264_videotoolbox -realtime 1 -profile:v 3 -level 51 -threads 4 -b:v 4000k -r 10 -map 0 -f segment -segment_time 05:00 screen_%03d.mkv -map "[out1]" -b:v 200k -f mpegts udp://127.0.0.1:1234

## UDP
ffmpeg -f avfoundation -framerate 30 -i "2" -f mpeg1video -b 200k -vf scale=380:-1  udp://127.0.0.1:1234
ffmpeg -f avfoundation -framerate 30 -i "2" -c:v h264_videotoolbox -realtime 1 -b:v 4000k -f mpegts udp://127.0.0.1:1234
ffmpeg -f avfoundation -framerate 30 -i "2" -vf scale=640:-1 -c:v h264_videotoolbox -realtime 1 -b:v 4000k -f mpegts udp://127.0.0.1:1234
ffplay udp://@127.0.0.1:1234
ffplay -fflags nobuffer -flags low_delay udp://127.0.0.1:1234

## RTP
#ffmpeg -f avfoundation -framerate 30 -i "2" -b 200k -vf scale=380:-1 -f rtp rtp://127.0.0.1:1234

## RTSP
-c:v copy

## H264
ffmpeg -h encoder=h264_videotoolbox
ffmpeg -t 20 -f avfoundation -i "2" -c:v h264_videotoolbox -realtime 1 -profile:v 3 -level 51 -b:v 4000k output.mkv
ffmpeg -f avfoundation -i "2" -c:v libx264 -preset ultrafast -qp 0 output.mkv
ffmpeg -f avfoundation -i "2" -c:v libx264 -preset medium -qp 0 -t 20 output.mkv -t 20

## H265
ffmpeg -h encoder=hevc_videotoolbox
ffmpeg -t 20 -f avfoundation -framerate 30 -video_size 640x480 -i "default" -c:v hevc_videotoolbox -tag:v hvc1 output.mkv

## time
mdfind .ttf
https://man7.org/linux/man-pages/man3/strftime.3.html
ffmpeg -f avfoundation -framerate 30 -i "2" -vf "drawtext=fontfile=arial.ttf:fontsize=24:fontcolor=yellow:text='%{localtime\:%a %d %b %Y %T}': box=1:boxcolor=black@0.8: x=7: y=7" -t 20 out.mkv
ffmpeg -f avfoundation -framerate 30 -i "2" -vf "drawtext=fontfile=arial.ttf:fontsize=24:fontcolor=yellow:text='%{pts\:gmtime\:0\:%M\\\\\:%S}': box=1:boxcolor=black@0.8: x=7: y=7" -t 20 out.mkv
ffmpeg -f avfoundation -framerate 30 -i "2" -vf "drawtext=fontfile=arial.ttf:fontsize=24:fontcolor=yellow:text='%{localtime\:%a %d %b %Y %T} %{pts\:gmtime\:0\:%M\\\\\:%S}': box=1:boxcolor=black@0.5: x=7: y=7" -t 20 out.mkv

## record sound
ffmpeg -f avfoundation -i ":0" -codec:a copy output.wav

### output file to 24 fps
-r 24 
 
ffmpeg -devices
-hide_banner
-t duration
-fs limit_size
-metadata title="my title"
-target svcd
-timelimit duration (global)
b:v 64k -bufsize 64k
-f mpegts udp://10.1.0.102:1234

ffmpeg -f avfoundation -i "1:0" -y -framerate 30 Screen.mov

### Record video from video device 0 and audio from audio device 0 into out.avi
ffmpeg -f avfoundation -framerate 30 -i "0" out.avi 
ffmpeg -f avfoundation -framerate 30 -i "0" out.mpg 
ffmpeg -f avfoundation -framerate 30 -i "0" out.mov 
ffmpeg -f avfoundation -framerate 30 -i "1" out.avi 
ffmpeg -f avfoundation -framerate 30 -i "1" out.flv 
ffmpeg -f avfoundation -framerate 30 -i "1" out.mkv 
ffmpeg -f avfoundation -framerate 30 -i "1" -r 10 out.mp4
ffmpeg -f avfoundation -pix_fmt uyvy422 -framerate 30 -i "1" output.mkv

ffmpeg -f avfoundation -pixel_format bgr0 -framerate 30 -i "default:none" out.avi

ffmpeg -f v4l2 -i /dev/video0 -vcodec rawvideo -delay 20 -f mpegts udp://127.0.0.1:1234

### segments
(1)  ffmpeg -i rtsp://<ip>/stream1 -codec copy -map 0 -f segment
-segment_time 4 -segment_format mp4 segment%d.m4s

This creates separate mp4 files for every 4 seconds of video, but each file
is stand-alone... i.e. each one has it's own 'moov' box... they aren't
fragments of the same stream, but rather independent streams.

-f segment -segment_time 4 -reset_timestamps 1

how to open udp: channel
how to mosaic two stream into one
how to segment files
write overlay inside video
multiple output

https://stackoverflow.com/questions/63460919/how-to-improve-the-output-video-quality-with-ffmpeg-and-h264-videotoolbox-flag

https://windpoly.run/posts/docker-ipc-fifo/

ffmpeg -y -i out.mp4 -r 60 -vf "minterpolate=fps=30:mi_mode=mci:mc_mode=aobmc:me_mode=bidir:vsbmc=1" -filter:v "setpts=(2/(2+0.1*T))*PTS" new.mp4