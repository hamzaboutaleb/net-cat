package utils

import "net"

func PrintLogo(c net.Conn) {
	c.Write([]byte("Welcome to TCP-Chat!\n"))
	c.Write([]byte("         _nnnn_\n"))
	c.Write([]byte("        dGGGGMMb\n"))
	c.Write([]byte("       @p~qp~~qMb\n"))
	c.Write([]byte("       M|@||@) M|\n"))
	c.Write([]byte("       @,----.JM|\n"))
	c.Write([]byte("      JS^\\__/  qKL\n"))
	c.Write([]byte("     dZP        qKRb\n"))
	c.Write([]byte("    dZP          qKKb\n"))
	c.Write([]byte("   fZP            SMMb\n"))
	c.Write([]byte("   HZM            MMMM\n"))
	c.Write([]byte("   FqM            MMMM\n"))
	c.Write([]byte(" __| \".        |\\dS\"qML\n"))
	c.Write([]byte(" |    `.       | `' \\Zq\n"))
	c.Write([]byte("_)      \\.___.,|     .'\n"))
	c.Write([]byte("\\____   )MMMMMP|   .'\n"))
	c.Write([]byte("     `-'       `--'\n"))
	c.Write([]byte("[ENTER YOUR NAME]:"))
}
