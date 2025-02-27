package hookApiMgr_test

import (
	"github.com/ddkwork/hyperdbgui/src/fnTable/hookApiMgr"
	"github.com/ddkwork/librarygo/src/mylog"
	"github.com/ddkwork/librarygo/src/stream"
	"testing"
)

func TestName(t *testing.T) {
	h := hookApiMgr.New()
	h.TestIopXxxControlFile()
}

func TestName1(t *testing.T) {
	/*
	    	!cpuid script {
	   			 	printf("$pname\t%s\n", $pname);
	       		 	printf("rax,rbx,rcx,rdx: %llx\t%llx\t%llx\t%llx\t\n", @rax,@rbx,@rcx,@rdx);
	       		    if ($context == 1) {
	       		 	    @rax=0x000306A9;
	       		 		@rcx=0x3DBAE3BF;
	       		 		@rdx=0xBFEBFBFF;
	       		      }
	   		 	}

	*/
	/*
			 	if (
		db($pname  ) == 73 &&
		db($pname+1) == 73 &&
		db($pname+2) == 64 &&
		db($pname+3) == 2e &&
		db($pname+3) == 65 &&
		db($pname+3) == 78 &&
		db($pname+3) == 65
			) {

			    }
	*/
	ABCDE := stream.NewString("ABCDE")
	ssd := stream.NewString("ssd.exe")
	mylog.HexDump("ABCDE", ABCDE.Buffer.Bytes())
	mylog.HexDump("ssd.exe", ssd.Buffer.Bytes())

}
