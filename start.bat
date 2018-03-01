@echo off

set chain=e-coin
set output=.env
set multi_dir=%APPDATA%\MultiChain\%chain%

echo [INFO] (T-7) The chain %chain% will be created if not existing in the %multi_dir% directory

if not exist %CD%\multichain goto download

:multichain
cd multichain
if not exist %multi_dir% goto initialize

:start
start multichaind.exe %chain%
echo [OK] (T-4) Multichain daemon has been launched on a new windows (don't close it)
cd ..

if not exist %output% goto configuration

:compile
if not exist %CD%\frontend\dist goto front
:build
if not exist main.exe goto back
:launch
start main.exe
echo [FINISH] (T+1) The server is now launched! You can also launch it manually with ./main.exe
pause
exit

:download
echo [INFO] (T-6) Downloading MultiChain... Please wait...
powershell -Command "(New-Object Net.WebClient).DownloadFile('https://www.multichain.com/download/multichain-windows-2.0-alpha-2.zip', 'multichain.zip')"
powershell -Command "Expand-Archive multichain.zip"
del multichain.zip
echo [OK] Multichain successfully installed in ./multichain
goto multichain


:initialize
echo [INFO] (T-5) Initialize new Blockchain
multichain-util.exe create %chain%
echo [OK] The blockchain has been created and is ready for launch
goto start


:configuration
set conf=%multi_dir%\multichain.conf

for /f "delims== tokens=1,2" %%G in (%conf%) do set %%G=%%H
echo MULTICHAIN_CHAIN_NAME=%chain% >> %output%
echo MULTICHAIN_HOST=localhost >> %output%
echo MULTICHAIN_RPC_USER=%rpcuser% >> %output%
echo MULTICHAIN_RPC_PASSWORD=%rpcpassword% >> %output%

for /f "delims==  tokens=1,2" %%B in (%multi_dir%\params.dat) do set %%B=%%C
for /f "delims=#  tokens=1,2" %%E in ("%default-rpc-port %") do set port=%%E
echo MULTICHAIN_PORT=%port% >> %output%

echo [INFO] All Multichain parameters have been filled in %output% file

echo AUTH0_DOMAIN=lyonnister.eu.auth0.com >> %output%
echo AUTH0_AUDIENCE=https://e-coin.io >> %output%

echo [INFO] Auth0 demo parameters have been filled in %output% file
echo [OK] %output% is ready!
goto compile


:front
cd frontend
echo [INFO] (T-2) NPM install…
npm run install
echo [INFO] NPM build
npm run build
echo [OK] Front-end ready!
cd ..
goto build

:back
echo [INFO] (T-1) Build server...
go build main.go
echo [OK] You can now launch it with ./main.exe
goto launch