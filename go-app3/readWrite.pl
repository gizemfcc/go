open (in, "<", "results.txt") or die "Could not open";
open (out, "+>>primeResults.txt") or die "Cannot open primeResults.txt: $!";
while(<in>){
    if(index($_, '=> PRIME')> -1){
        $subs = substr $_, 0, index($_, '=> PRIME');
	    print "$subs\n";
	    print out "$subs\n"
    }
}




