# Reconcile

## Overview

Reconcile Internal System Transaction Report : Multiple Banks Transaction Report
#

## Features

- **Tools - Generate CSV** :
  - Generate CSV file for Internal System and (Multiple) Bank with different format
- **Reconcile**
  - Compare CSV files
  - Generate MisMatch file result : System & (Multiple Files) for Bank
#


## Usage
### Tools Generate CSV
```
make tools-generate-csv   
```

### Reconcile
```
make reconcile system=/home/opannapo/PROJECT-CODE/PRIB/GIT/reconcile/sample-data/system bank=/home/opannapo/PROJECT-CODE/PRIB/GIT/reconcile/sample-data/bank start=2023-01-01 end=2025-03-11
```

### Reproduce Step
**1. Generate CSV using tools CLI app**
- Output
   - Path `/reconcile/sample-data/system/SYSTEM.csv`
   - Path :
     - `/reconcile/sample-data/bank/BCA.csv`
     - `/reconcile/sample-data/bank/BRI.csv`
     - `/reconcile/sample-data/bank/MANDIRI.csv`

**2. Modify Bank File**
- Path :
     - `/reconcile/sample-data/bank/BCA.csv`
     - `/reconcile/sample-data/bank/BRI.csv`
       
**3. Run Reconcile CLI Apps**
- Output
  - `/reconcile/sample-data/mismatches/MISMATCHES-BCA.csv.csv`
  - `/reconcile/sample-data/mismatches/MISMATCHES-BRI.csv.csv`
  - `/reconcile/sample-data/mismatches/MISMATCHES-SYSTEM.csv`



#
# Demo
[Screencast Output - Tools Generate & Generate Mismatch .webm](https://github.com/user-attachments/assets/9089edc8-c19b-4fea-bcfc-4ba03d104971)


## Specs of the Development Laptop
```bash

neofetch

            .-/+oossssoo+/-.               @legion 
        `:+ssssssssssssssssss+:`           --------------- 
      -+ssssssssssssssssssyyssss+-         OS: Ubuntu 22.04.5 LTS x86_64 
    .ossssssssssssssssssdMMMNysssso.       Host: 82RB Legion 5 15IAH7H 
   /ssssssssssshdmmNNmmyNMMMMhssssss/      Kernel: 6.8.0-51-generic 
  +ssssssssshmydMMMMMMMNddddyssssssss+     Resolution: 2560x1440 
 /sssssssshNMMMyhhyyyyhmNMMMNhssssssss/    DE: GNOME 42.9 
.ssssssssdMMMNhsssssssssshNMMMdssssssss.   WM Theme: Adwaita 
+sssshhhyNMMNyssssssssssssyNMMMysssssss+   Theme: Yaru-bark-dark [GTK2/3] 
ossyNMMMNyMMhsssssssssssssshmmmhssssssso   CPU: 12th Gen Intel i7-12700H (20) @ 4.600GHz 
ossyNMMMNyMMhsssssssssssssshmmmhssssssso   GPU: NVIDIA GeForce RTX 3060 Mobile / Max-Q 
+sssshhhyNMMNyssssssssssssyNMMMysssssss+   GPU: Intel Alder Lake-P 
.ssssssssdMMMNhsssssssssshNMMMdssssssss.   Memory: 9276MiB / 15714MiB 
 /sssssssshNMMMyhhyyyyhdNMMMNhssssssss/     
  +sssssssssdmydMMMMMMMMddddyssssssss+      
   /ssssssssssshdmNNNNmyNMMMMhssssss/       
    .ossssssssssssssssssdMMMNysssso.        
      -+sssssssssssssssssyyyssss+-          
        `:+ssssssssssssssssss+:`            
            .-/+oossssoo+/-.
                                                                   
                                                                   
```

