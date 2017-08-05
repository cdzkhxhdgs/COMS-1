# COMS-1
Tools for COMS-1 LRIT satellite data. Requires Python 3.

| Component     | Description   |
| ------------- | ------------- |
| [lrit-header.py](#lrit-headerpy)  | Parses LRIT file and displays header information in a human-readable format.  |
| [lrit-additional.py](#lrit-additionalpy)  | Extracts data from LRIT Additional Data (ADD) files.  |
| [coms.py](coms.py) | Common functions and variables. |

## lrit-header.py
Parses LRIT file and displays header information in a human-readable format.
```
usage: lrit-header.py [-h] PATH

Parses LRIT file and displays header information in a human-readable format.

positional arguments:
  PATH        Input LRIT file

optional arguments:
  -h, --help  show this help message and exit
```

### Sample output
```
python3.6 lrit-header.py samples/lrit/IMG_ENH_01_IR1_20120101_000920_01.lrit
[Type 000 : Offset 0x0000] Primary Header:
	Header length:         16 (0x10)
	File type:             0, Image data (IMG)
	Total header length:   4972 (0x136C)
	Data length:           3824184 (0x3A5A38)

[Type 001 : Offset 0x0010] Image Structure Header:
	Header length:         9 (0x9)
	Bits per pixel:        8
	Image:                 Extended Northern Hemisphere (ENH)
	  - Columns: 1547
	  - Lines:   309
	Compression:           0, None

[Type 002 : Offset 0x0019] Image Navigation Header:
	Header length:         51 (0x33)
	Projection:            Normalized Geostationary Projection (GEOS)
	Longitude:             128.2° E
	Column scaling factor: 8170135
	Line scaling factor:   4286797161
	Column offset:         773
	Line offset:           1010

[Type 003 : Offset 0x004C] Image Data Function Header:
	Header length:         4810 (0x12CA)
	Data Definition Block:
	  - dumped to "samples/lrit/IMG_ENH_01_IR1_20120101_000920_01_IDF-DDB.txt"

[Type 004 : Offset 0x1316] Annotation Header:
	Header length:         41 (0x29)
	Text data:             "IMG_ENH_01_IR1_20120101_000920_01.lrit"

[Type 005 : Offset 0x133F] Time Stamp Header:
	Header length:         10 (0xA)
	P Field:               01000000
	  - Extension flag:    0 (No extension)
	  - Time code ID:      100 (1958 January 1 epoch - Level 1 Time Code)
	  - Detail bits:       0000
	T Field:               010011010000101000000101000110001110111010000000
	  - Day counter:       19722 (01/01/2012 - DD/MM/YYYY)
	  - Milliseconds:      85520000 (23:45:20 - HH:MM:SS)

[Type 007 : Offset 0x1349] Key Header:
	Header length:         7 (0x7)
	Encryption key:        0x0 (disabled)

[Type 128 : Offset 0x1350] Image Segmentation Information Header:
	Header length:         7 (0x7)
	Segment number:        1 of 4
	Line num of image:     1
```

## lrit-additional.py
Extracts data from LRIT Additional Data (ADD) files. Data includes Alpha-numeric text (ANT), CMDPS (CT/CTT/CTH), and GOCI.

```
usage: lrit-additional.py [-h] PATH

Extracts data from LRIT Additional Data (ADD) files. Data includes Alpha-
numeric text (ANT), CMDPS (CT/CTT/CTH), and GOCI.

positional arguments:
  PATH        Input LRIT file

optional arguments:
  -h, --help  show this help message and exit
```

### Sample output (Alpha-numeric Text)
```
python3.6 lrit-additional.py samples/lrit/ADD_ANT_01_20120101_113500_00.lrit
[Type 0 : Offset 0x0] Primary Header:
	Header length:         16 (0x10)
	File type:             2, Alpha-numeric text (ANT)
	Total header length:   70 (0x46)
	Data length:           78088 (0x13108)

[Type 4 : Offset 0x10] Annotation Header:
	Header length:         37 (0x25)
	Text data:             "ADD_ANT_01_20120101_113500_00.lrit"

[Type 5 : Offset 0x35] Time Stamp Header:
	Header length:         10 (0xA)
	P Field:               01000000
	  - Extension flag:    0 (No extension)
	  - Time code ID:      100 (1958 January 1 epoch - Level 1 Time Code)
	  - Detail bits:       0000
	T Field:               010011010000101100000001001110001010011010000011
	  - Day counter:       19723 (02/01/2012 - DD/MM/YYYY)
	  - Milliseconds:      20489859 (05:41:29 - HH:MM:SS)

[Type 7 : Offset 0x3F] Key Header:
	Header length:         7 (0x7)
	Encryption key:        0x0 (disabled)

Additional Data dumped to "samples/lrit/ADD_ANT_01_20120101_113500_00_DATA.txt"
```

### Sample output (CMDPS - CTH)
```
python3.6 lrit-additional.py samples/lrit/ADD_CTH_02_20120101_033200_00.lrit
[Type 0 : Offset 0x0] Primary Header:
	Header length:         16 (0x10)
	File type:             128, COMS Meteorological Data Processing System (CMDPS) analysis data
	Total header length:   70 (0x46)
	Data length:           4689808 (0x478F90)

[Type 4 : Offset 0x10] Annotation Header:
	Header length:         37 (0x25)
	Text data:             "ADD_CTH_02_20120101_033200_00.lrit"

[Type 5 : Offset 0x35] Time Stamp Header:
	Header length:         10 (0xA)
	P Field:               01000000
	  - Extension flag:    0 (No extension)
	  - Time code ID:      100 (1958 January 1 epoch - Level 1 Time Code)
	  - Detail bits:       0000
	T Field:               010011010000101100000000101010101101100101101111
	  - Day counter:       19723 (02/01/2012 - DD/MM/YYYY)
	  - Milliseconds:      11196783 (03:06:36 - HH:MM:SS)

[Type 7 : Offset 0x3F] Key Header:
	Header length:         7 (0x7)
	Encryption key:        0x0 (disabled)

Additional Data dumped to "samples/lrit/ADD_CTH_02_20120101_033200_00_DATA.png"
```

## Sample data
LRIT and HRIT sample data was obtained from [Korea Meteorological Administration's (KMA) National Meteorological Satellite Center (NMSC)](http://nmsc.kma.go.kr/html/homepage/en/chollian/Introduction/selectIntroduction.do). Code examples and xRIT Mission Specific Implementation documents are also provided.

### VIS Image
![MI VIS 23.07.17 0515UTC](https://raw.githubusercontent.com/sam210723/COMS-1/master/samples/coms_mi_le1b_vis_cf_201707230515.png)
Image obtained from [KMA NMSC](http://nmsc.kma.go.kr/html/homepage/en/satellite/searchSatelliteImageN.do?data_type=1001)

### Additional Data (ADD)
Additional Data types:
 * [Alpha-numeric Text (ANT)](#alpha-numeric-text-ant)
 * [COMS Meteorological Data Processing System (CMDPS) analysis data](#cmdps)
 * Numerical Weather Prediction (NWP) data
 * Geostationary Ocean Color Imager (GOCI) data
 * KMA typhoon information

#### Alpha-numeric Text (ANT)
```
COMS MI LRIT WOP(Weekly Operation Plan) 

BY NATIONAL METEOROLOGICAL SATELLITE CENTER, KOREA METEROLOGICAL ADMINSTRATION

DISSEMINATION SCHEDULE FROM 2012/01/02 TO 2012/01/08
UPDATE 2012/01/01

NOTE: 

＊ FOR SEMI REAL-TIME INFORMATIN ON COMS OPERATION, PLEASE REFER TO FOLLOWING WEBSITE. (http://nmsc.kma.go.kr)

＊An wheel off-loading maneuver is scheduled from [00:45:00.000] on [2012/01/02~01/08]. 
   For about 20 minutes, the image quality might be temporarily deteriorated. 
   For about 20 minutes, the image might be disseminated with chipping. (See the following schedule for the details)
＊An wheel off-loading maneuver is scheduled from [06:45:00.000] on [2012/01/02~01/08]. 
   For about 20 minutes, the image quality might be temporarily deteriorated. 
   For about 20 minutes, the image might be disseminated with chipping. (See the following schedule for the details)

SCHEDULE:

TIME(UTC)        LRIT    ABBR_ID 02 03 04 05 06 07 08
0009-0012        E01        ENH  O  O  O  O  O  O  O
0024-0027        E02        ENH  O  O  O  O  O  O  O
0028-0030        C01         CT  O  O  O  O  O  O  O
0032-0034        C01        CTH  O  O  O  O  O  O  O
...
```
Alpha-numeric Text sample continued in [samples/lrit/ADD_ANT_01_20120101_113500_00_DATA.txt](samples/lrit/ADD_ANT_01_20120101_113500_00_DATA.txt)

#### CMDPS
Cloud Type (CT) | Cloud Top Height (CTH) | Cloud Top Temperature (CTT)
------------ | ------------- | -------------
![Cloud Type (CT)](samples/lrit/ADD_CT_02_20120101_032800_00_DATA.png) | ![Cloud Top Height (CTH)](samples/lrit/ADD_CTH_02_20120101_033200_00_DATA.png) | ![Cloud Top Temperature (CTT)](samples/lrit/ADD_CTT_02_20120101_033500_00_DATA.png)
