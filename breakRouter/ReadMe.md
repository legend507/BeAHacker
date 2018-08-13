# Use known vulnerabilities to break routers
1. www.exploit-db.com, a good start point to search for known vulnerabilities
2. Google "xxx(router name) vulns", I may find more known vulnerabilities
(The vulns can be very different)
3. for some vulns, there is a chance that I can get a binary config file from a router
4. find a xxx (maker) config file decompressor to find username:password to login to the router

(use nmap)
5. goto https://nmap.org/nsedoc/, to download the scripts (scripts are used to exploit some known vulns); how to download is basically copy the code, touch a new file on my PC, and paste
6. goto /usr/share/nmap/scripts/, to check all .nse scripts I have (under ./nselib/, are other .lua scripts)
7. "nmap -p80 --script http-tplink-dir-traversal.nse --script-args rfile=/tmp/ath0.ap_bss -d -n -Pn 192.168.0.1", to use the script
(Still don't know what the script does, need to try on my Buffalo router.)
After this, I can use some other tools (e.g., Reaver) to do further hacking.
