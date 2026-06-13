# Python

## environment variables

- `PYTHONPATH`, for Python to lookup modules. Usually, it looks for `site-packages` folder.
Example: `export PYTHONPATH="/e/writable/Python38/site-packages:/c/Program\ Files/Python38/Lib/site-packages:/c/Users/tliu/AppData/Roaming/Python/Python38/site-packages"

- `PYTHONUSERBASE`, if `--uesr` is given, it will look given folder to install modules.
Example: `export PYTHONUSERBASE="E:\\wirtable"`, it will install modules into `E:\writable\Python38\Scripts` and `E:\writable\Python38\site-packages`(later one can be specified in to PYTHONPATH ahead) for `python setup.py install --user`

- `PYTHONHOME`, the location of the standard Python libraries
