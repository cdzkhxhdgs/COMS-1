# COMS-1
Tools for COMS-1 satellite data

## lrit-header.py
Sample output using ``samples/lrit/IMG_ENH_01_IR1_20120101_000920_01.lrit``
```
[Type 0 : Offset 0x0] Primary Header:
	Header length:         16
	File type:             0, Image data (IMG)
	Total header length:   4972 (0x136C)
	Data length:           3824184 (0x3A5A38)

[Type 1 : Offset 0x10] Image Structure Header:
	Header length:         9
	Bits per pixel:        8
	Image:                 Extended Northern Hemisphere (ENH)
	  - Columns: 1547
	  - Lines:   309
	Compression:           0, None

[Type 2 : Offset 0x19] Image Navigation Header:
	Header length:         51
	Projection:            Normalized Geostationary Projection (GEOS)
	Longitude:             128.2° E
	Column scaling factor: 8170135
	Line scaling factor:   4286797161
	Column offset:         773
	Line offset:           1010

[Type 3 : Offset 0x4C] Image Data Function Header:
	Header length:         4810 (0x12CA)
	Data Definition Block:
	  - dumped to "samples/lrit/IMG_ENH_01_IR1_20120101_000920_01_IDF-DDB.txt"

[Type 4 : Offset 0x1316] Annotation Header:
	Header length:         41 (0x29)
	Text data:             "IMG_ENH_01_IR1_20120101_000920_01.lrit"

[Type 5 : Offset 0x133F] Time Stamp Header:
	Header length:         10
	P Field:               01000000
	  - Extension flag:    0 (No extension)
	  - Time code ID:      100 (1958 January 1 epoch - Level 1 Time Code)
	  - Detail bits:       0000
	T Field:               010011010000101000000101000110001110111010000000
	  - Day counter:       19722 (01/01/2012 - DD/MM/YYYY)
	  - Milliseconds:      85520000 (23:45:20 - HH:MM:SS)

[Type 7 : Offset 0x1349] Key Header:
	Header length:         7
	Encryption key:        0x0 (disabled)

[Type 128 : Offset 0x1350] Image Segment Definition Header:
	Header length:         7
	Segment number:        1 of 4
	Line num of segment:   1
```

## Sample data
LRIT and HRIT sample data was obtained from [Korea Meteorological Administration's (KMA) National Meteorological Satellite Center (NMSC)](http://nmsc.kma.go.kr/html/homepage/en/chollian/Introduction/selectIntroduction.do). Code examples and xRIT Mission Specific Implementation documents are also provided.

## Sample VIS Image
![MI VIS 23.07.17 0515UTC](https://raw.githubusercontent.com/sam210723/COMS-1/master/samples/coms_mi_le1b_vis_cf_201707230515.png)
Image obtained from [KMA NMSC](http://nmsc.kma.go.kr/html/homepage/en/satellite/searchSatelliteImageN.do?data_type=1001)
