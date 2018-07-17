(string) (len=5889) "IDENT\tSUBROUTINE\nIDENT\tDDIAPA\n(\t(\nIDENT\tIX\n,\t,\nIDENT\tIY\n)\t)\nCOMMENT\tC     (Add to Path)\nCOMMENT\tC     Add the point (IX,IY) to the current path, which is assumed\nCOMMENT\tC     to have  already been  started by  a call  to DDIBPA.   The\nCOMMENT\tC     current point and visibility are updated in COMMON.\nCOMMENT\tC     (27-May-1991)\nCOMMENT\tC- - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -\nCOMMENT\tC\nCOMMENT\tC     EXTERNAL REFERENCES (FUNCTION,SUBROUTINE,COMMON)\nCOMMENT\tC\nCOMMENT\tC     EXTERNAL REFS       TKIOLL,      TKIOWB,      TKIOWH,      TKIOWN\nCOMMENT\tC\nIDENT\tINTEGER\nIDENT\tTKIOLL\nCOMMENT\tC- - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -\nCOMMENT\tC\nCOMMENT\tC     NON-COMMON VARIABLES\nCOMMENT\tC\nCOMMENT\tC     SETS                BLANK\nCOMMENT\tC\nIDENT\tINTEGER\nIDENT\tBLANK\n,\t,\nIDENT\tDUMMY\n,\t,\nIDENT\tIX\n,\t,\nIDENT\tIY\nCOMMENT\tC\nCOMMENT\tC-----------------------------------------------------------------------\nCOMMENT\tC                         P o s t S c r i p t\nCOMMENT\tC           D i s p l a y   D e v i c e   I n t e r f a c e\nCOMMENT\tC                       C O M M O N   B l o c k\nCOMMENT\tC\nCOMMENT\tC     DSSIZE                   display surface size in cm\nCOMMENT\tC     IMAGE                    image transformation in effect\nCOMMENT\tC     (LASTX,LASTY,VILAST)     last position and visibility of pen\nCOMMENT\tC     LINEIN                   line intensity (0..1 scale)\nCOMMENT\tC     LINEWT                   line weight (1..25 scale)\nCOMMENT\tC     MAGFAC                   resolution magnification factor\nCOMMENT\tC     (MOVEX,MOVEY)            position of first point in path\nCOMMENT\tC     NPATH                    number of line segments in current page\nCOMMENT\tC     PLSTEP                   plotter step size in cm\nCOMMENT\tC     (PSXOFF,PSYOFF)          image offset in PostScript units (bp)\nCOMMENT\tC     ROTATE                   plot frame is rotated\nCOMMENT\tC     (SX,SY)                  scale factors (unit square to plot steps)\nCOMMENT\tC     TIMAGE(*,*)              image transformation matrix\nCOMMENT\tC     (XMAX,YMAX,ZMAX)         normalized device space extents\nCOMMENT\tC     (MAXX,MAXY,MINX,MINY)    actual coordinate limits reached\nCOMMENT\tC\nCOMMENT\tC-----------------------------------------------------------------------\nIDENT\tINTEGER\nIDENT\tLASTX\n,\t,\nIDENT\tLASTY\n,\t,\nIDENT\tLINEWT\n,\t,\nIDENT\tMAGFAC\nIDENT\tINTEGER\nIDENT\tMAXX\n,\t,\nIDENT\tMAXY\n,\t,\nIDENT\tMINX\n,\t,\nIDENT\tMINY\nIDENT\tINTEGER\nIDENT\tMOVEX\n,\t,\nIDENT\tMOVEY\n,\t,\nIDENT\tNPATH\n,\t,\nIDENT\tPSXOFF\nIDENT\tINTEGER\nIDENT\tPSYOFF\nCOMMENT\tC\nIDENT\tLOGICAL\nIDENT\tDVINIT\n,\t,\nIDENT\tIMAGE\n,\t,\nIDENT\tROTATE\n,\t,\nIDENT\tVILAST\nCOMMENT\tC\nIDENT\tREAL\nIDENT\tDSSIZE\n,\t,\nIDENT\tLINEIN\n,\t,\nIDENT\tPLSTEP\n,\t,\nIDENT\tSX\nIDENT\tREAL\nIDENT\tSY\n,\t,\nIDENT\tTIMAGE\n,\t,\nIDENT\tXMAX\n,\t,\nIDENT\tYMAX\nIDENT\tREAL\nIDENT\tZMAX\nCOMMENT\tC\nIDENT\tCOMMON\n/\t/\nIDENT\tDDI01\n/\t/\nIDENT\tDSSIZE\n,\t,\nIDENT\tLINEIN\n,\t,\nIDENT\tPLSTEP\n,\t,\nIDENT\tSX\nIDENT\tCOMMON\n/\t/\nIDENT\tDDI01\n/\t/\nIDENT\tSY\n,\t,\nIDENT\tTIMAGE\n(\t(\nIDENT\t4\n,\t,\nIDENT\t4\n)\t)\n,\t,\nIDENT\tXMAX\n,\t,\nIDENT\tYMAX\nIDENT\tCOMMON\n/\t/\nIDENT\tDDI01\n/\t/\nIDENT\tZMAX\n,\t,\nIDENT\tLASTX\n,\t,\nIDENT\tLASTY\n,\t,\nIDENT\tLINEWT\nIDENT\tCOMMON\n/\t/\nIDENT\tDDI01\n/\t/\nIDENT\tMAGFAC\n,\t,\nIDENT\tMAXX\n,\t,\nIDENT\tMAXY\n,\t,\nIDENT\tMINX\nIDENT\tCOMMON\n/\t/\nIDENT\tDDI01\n/\t/\nIDENT\tMINY\n,\t,\nIDENT\tMOVEX\n,\t,\nIDENT\tMOVEY\n,\t,\nIDENT\tNPATH\nIDENT\tCOMMON\n/\t/\nIDENT\tDDI01\n/\t/\nIDENT\tPSXOFF\n,\t,\nIDENT\tPSYOFF\n,\t,\nIDENT\tDVINIT\n,\t,\nIDENT\tIMAGE\nIDENT\tCOMMON\n/\t/\nIDENT\tDDI01\n/\t/\nIDENT\tROTATE\n,\t,\nIDENT\tVILAST\nCOMMENT\tC\nCOMMENT\tC\nIDENT\tDATA\nIDENT\tBLANK\n/\t/\nIDENT\t32\n/\t/\nCOMMENT\tC\nCOMMENT\tC     We can generate an  entry of size \"-ddddd -ddddd R \"; try to\nCOMMENT\tC     keep line lengths under 80 characters.\nCOMMENT\tC\nIDENT\tIF\n(\t(\nIDENT\tTKIOLL\n(\t(\nIDENT\tDUMMY\n)\t)\n.\t.\nIDENT\tGT\n.\t.\nIDENT\t64\n)\t)\nIDENT\tCALL\nIDENT\tTKIOWH\n(\t(\nIDENT\t2\nIDENT\tH\nIDENT\tN\n,\t,\nIDENT\t2\n)\t)\nCOMMENT\tC\nCOMMENT\tC     Use relative coordinates for lines to reduce the  number of\nCOMMENT\tC     bytes in the  verbose  PostScript  command language.   This\nCOMMENT\tC     shortens the plot file by 5 to 10 percent.\nCOMMENT\tC\nCOMMENT\tC     We optimize output into 3 different forms:\nCOMMENT\tC\nCOMMENT\tC     # # R   (relative lineto)\nCOMMENT\tC     # X     (relative lineto with delta-y = 0)\nCOMMENT\tC     # Y     (relative lineto with delta-x = 0)\nCOMMENT\tC\nCOMMENT\tC     The  latter two cases  occur frequently enough  to be worth\nCOMMENT\tC     taking advantage of  to compress the  output and reduce the\nCOMMENT\tC     parsing  time.    We cannot  discard zero length  segments,\nCOMMENT\tC     however, because that  would discard single  points from  a\nCOMMENT\tC     point plot.\nCOMMENT\tC\nIDENT\tIF\n(\t(\nIDENT\tIX\n.\t.\nIDENT\tEQ\n.\t.\nIDENT\tLASTX\n)\t)\nIDENT\tGO\nIDENT\tTO\nIDENT\t10\nIDENT\tIF\n(\t(\nIDENT\tIY\n.\t.\nIDENT\tEQ\n.\t.\nIDENT\tLASTY\n)\t)\nIDENT\tGO\nIDENT\tTO\nIDENT\t20\nIDENT\tCALL\nIDENT\tTKIOWN\n(\t(\nIDENT\tIX\n-\t-\nIDENT\tLASTX\n)\t)\nIDENT\tCALL\nIDENT\tTKIOWB\n(\t(\nIDENT\tBLANK\n)\t)\nIDENT\tCALL\nIDENT\tTKIOWN\n(\t(\nIDENT\tIY\n-\t-\nIDENT\tLASTY\n)\t)\nIDENT\tCALL\nIDENT\tTKIOWH\n(\t(\nIDENT\t3\nIDENT\tH\nIDENT\tR\n,\t,\nIDENT\t3\n)\t)\nIDENT\tGO\nIDENT\tTO\nIDENT\t30\nIDENT\t10\nIDENT\tCALL\nIDENT\tTKIOWN\n(\t(\nIDENT\tIY\n-\t-\nIDENT\tLASTY\n)\t)\nIDENT\tCALL\nIDENT\tTKIOWH\n(\t(\nIDENT\t3\nIDENT\tH\nIDENT\tY\n,\t,\nIDENT\t3\n)\t)\nIDENT\tGO\nIDENT\tTO\nIDENT\t30\nIDENT\t20\nIDENT\tCALL\nIDENT\tTKIOWN\n(\t(\nIDENT\tIX\n-\t-\nIDENT\tLASTX\n)\t)\nIDENT\tCALL\nIDENT\tTKIOWH\n(\t(\nIDENT\t3\nIDENT\tH\nIDENT\tX\n,\t,\nIDENT\t3\n)\t)\nIDENT\t30\nIDENT\tNPATH\n=\t=\nIDENT\tNPATH\n+\t+\nIDENT\t1\nIDENT\tLASTX\n=\t=\nIDENT\tIX\nIDENT\tLASTY\n=\t=\nIDENT\tIY\nIDENT\tVILAST\n=\t=\n.\t.\nIDENT\tTRUE\n.\t.\nIDENT\tMAXX\n=\t=\nIDENT\tMAX0\n(\t(\nIDENT\tMAXX\n,\t,\nIDENT\tIX\n)\t)\nIDENT\tMINX\n=\t=\nIDENT\tMIN0\n(\t(\nIDENT\tMINX\n,\t,\nIDENT\tIX\n)\t)\nIDENT\tMAXY\n=\t=\nIDENT\tMAX0\n(\t(\nIDENT\tMAXY\n,\t,\nIDENT\tIY\n)\t)\nIDENT\tMINY\n=\t=\nIDENT\tMIN0\n(\t(\nIDENT\tMINY\n,\t,\nIDENT\tIY\n)\t)\nCOMMENT\tC\nCOMMENT\tC     Insert a  line break  periodically to  prevent  excessively\nCOMMENT\tC     long lines.\nCOMMENT\tC\nCOMMENT\tC     IF (MOD(NPATH,10) .EQ. 0) CALL TKIOWH (2H$N,2)\nCOMMENT\tC\nIDENT\t40\nIDENT\tRETURN\nIDENT\tEND\n"
