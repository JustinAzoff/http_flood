import flash.net.URLStream;
import flash.net.URLRequest;
import flash.events.Event;
import flash.events.ProgressEvent;
import flash.utils.ByteArray;
import flash.utils.Timer;

import flash.external.ExternalInterface;

class Flashflood
{
    static function main()
    {
        ExternalInterface.addCallback("start", start);
        trace("Ready");
    }
    static function start(seconds)
    {
        var st = new SpeedTester();
        st.test(seconds);
    }

}

class SpeedTester
{
    private var stream:URLStream;
    private var ba:ByteArray;
    private var bytesread:UInt;
    private var start_time:Date;
    private var stop_time:Date;

    public function new()
    {
        bytesread = 0;
        stream = new URLStream();
        ba = new ByteArray();
        stream.addEventListener(Event.OPEN, openHandler);
        stream.addEventListener(ProgressEvent.PROGRESS, progressHandler);
        stream.addEventListener(Event.COMPLETE, completeHandler);
    }

    public function test(seconds)
    {
        var request:URLRequest = new URLRequest("/flood?s=" + seconds);
        stream.load(request);
    }

    private function openHandler(event:Event)
    {
        trace("Starting download");

        start_time = Date.now();
    }
    private function progressHandler(event:Event)
    {
        while( stream.bytesAvailable != 0) {
           this.bytesread += stream.bytesAvailable;
           stream.readBytes(ba, stream.bytesAvailable);
        }
    }
    private function completeHandler(event:Event)
    {
        stop_time = Date.now();
        var elapsed = (stop_time.getTime() - start_time.getTime())/1000;
        var bps = this.bytesread / elapsed;
        var MBps = Math.round(100*bps/1024/1024)/100;
        trace("Read " + this.bytesread + " bytes in " + elapsed + ". " + MBps + " MB/s");
        ExternalInterface.call("result", MBps);
    }
}
