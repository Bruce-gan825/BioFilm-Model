# BioFilm-Model
 This is the final project for 02601.
The model is based on the biofilm model from https://pubmed.ncbi.nlm.nih.gov/29593289/ and https://pubs.acs.org/doi/10.1021/sb300031n. The simulation provides a graphical user interface written entirely in Golang that allows researchers to visualize and study the interactions of bacteria in the early stages of biofilm development 

## Installation and Setup 
!!! This project uses Go modules! To ensure that you can run the program smoothly on your machine, please ensure that GO111MODULE is set to on. If your environment does not have GO111MODULE set to on, you can enable Go modules for all projects within the GOPATH using the following command: 

set GO111MODULE=on

To run our biofilm simulator on your machine, make a clone of this repository: 

git clone https://github.com/Bruce-gan825/BioFilm-Model

Ensure that your version of Golang has GIO installed. You can install GIO by navigating to the terminal within your Go editor (Goland, Visual Studio Code, vim, etc.) and running the following command:

go install gio
go mod tidy

Alternatively, you can use the following command:

go install gioui.org/cmd/gogio@latest
gio version 
go mod tidy

Once GIO has been installed onto your Go machine, the program should be ready to run. You can run the program by opening your system's Terminal or Powershell app and executing the following commands:

cd go/src/EarlyBiofilm 
./EarlyBiofilm <time> <growth rate> <cell size> <threshold> <nutrition map filepath> 

## Demo Run
Our project is designed such that you can 
