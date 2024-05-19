# db_go

Добавьте свой json файл с ключами из Google Firebase Console и укажите его в opt := option.WithCredentialsFile("название вашего файла"):

Firebase Console -> Project settings -> Service accounts -> Generate new private key

(´·‿·`)


*************************
To initialize the database, you need to add your key file to opt := option.WithCredentialsFile("your file")::

Firebase Console -> Project settings -> Service accounts -> Generate new private key - download json file
