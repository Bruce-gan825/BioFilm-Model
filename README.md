# BioFilm-Model
This is the final project for 02601.
 
The model is based on the biofilm model from https://pubmed.ncbi.nlm.nih.gov/29593289/ and https://pubs.acs.org/doi/10.1021/sb300031n. The simulation provides a graphical user interface written entirely in Golang that allows researchers to visualize and study the interactions of bacteria in the early stages of biofilm development.

## Demo Video
See the project in action here: https://www.youtube.com/watch?v=6MiCJpkqquc

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

*Time: a decimal number that represents the interval of time for a single step of the simulation. For best results, set equal to 1.

*Growth Rate: a decimal number that represents the amount that a spherical cell grows upon consuming nutrients within the simulation. For best results, set equal to 0.06-0.10. 

*Cell Size: a decimal number that represents the maximum size a cell can reach before performing binary fission. For best results, set equal to 10-20. 

*Threshold: a decimal number that represents the minimum amount of nutrients a cell needs before performing a single growth event. For best results, set equal to 5-10. 

*Nutrition Map Filepath: the name of the desired nutrient map to test. You can use one of the provided maps to test (gopher.txt, dendritic.txt). If you do not want to use a custom nutrient map and would prefer to use a uniform map instead, write **false** for this input. 

## Demo Run
Our project is designed such that you can place any text file representing a 2D array for a nutrient map into the folder "NutritionBoardInputs". The program can then take in your desired nutrient map and initialize the growth environment prior to allowing the user to spawn cells and further nutrients. If you instead have a .png or .jpg image for which you want to convert into a 2D array representing the nutrient map, you can use our script FluorescenceToMap.py, which you can execute using the following command from the Terminal or Powershell: 

python3 FluorescenceToMap.py <input_image_path> <output_file_path> 

To perform a test run you can run one of the sample nutrient maps included with our project:

./EggTimer 1 0.07 8 10 test_nutrient_elevation.txt

Try spawning cells in the simulation and observing the growth stages of the biofilm as well as the number of Colony Forming Units (n) over time!

## Simulation Features
*Integrated mouse input: use the mouse to interact with the simulation by spawning in new spherical cells or by spawning in new nutrients for your biofilms. Use the scroll wheel to zoom in and out of the simulation.

*Integrated keyboard input: use WASD keys to move around the simulation. Use spacebar to pause/play the simulation. 

*Adjust growth parameters: you can adjust the growthRate, cellSize, and threshold variables once executing the simulation as necessary. This allows the program to be tuned to different species of spherical bacteria.

*Biofilm growth and diffusion: biofilms will naturally expand outwards by consuming nearby nutrients. Once biofilms grow to a large enough size, they will perform a "budding-off" event where part of the biofilm detaches into the nearby environment to colonize another solid surface. 

*Biofilm discrimination: different species of bacteria (represented by different cell colours) can interact with each other by forming boundaries with either each other or solid surfaces. 

*Biofilm vortexing: independent biofilms are able to continually rotate within themselves, replicating the behaviour of biofilms vortex to continually cycle nutrients. 

*Quorum sensing: cells within every biofilm continually release signalling particles into their environment; these particles cause the cells within a biofilm to prefer to clump together and form clusters. 

## Contributors

This project was created in the Fall of 2023 by graduate students Brian Zhang, Linqi Zhang, Yaoyuan Gan, and Tina Ryu. 
