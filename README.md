# Api Documentation

- base url : https://tes-teknikal-backend.onrender.com

### Get All House Listings

- Path : `/house/`
- Method : `GET`
- Response :
    ```json
      [
        {
          "ID": 1,
          "CreatedAt": "2025-10-16T19:04:20.757849Z",
          "UpdatedAt": "2025-10-16T19:04:20.757849Z",
          "DeletedAt": null,
          "title": "Rumah Depok",
          "poster": "LJ Hook",
          "description": "Depok rumah, tingkat 2."
        },
        {
          "ID": 2,
          "CreatedAt": "2025-10-16T19:07:04.319033Z",
          "UpdatedAt": "2025-10-16T19:07:04.319033Z",
          "DeletedAt": null,
          "title": "Apartemen Mlati Asri",
          "poster": "Mlati Asri",
          "description": "Apartemen hunian keluarga, 3 juta/bulan."
        }
      ]
    ```

### Get House Listing

- Path : `/house/{id}`
- Method : `GET`
- Response :

   ```json
      {
        "ID": 1,
        "CreatedAt": "2025-10-16T19:04:20.757849Z",
        "UpdatedAt": "2025-10-16T19:04:20.757849Z",
        "DeletedAt": null,
        "title": "Rumah Depok",
        "poster": "LJ Hook",
        "description": "Depok rumah, tingkat 2."
      }
   ```

### Create House Listing

- Path : `/house/`
- Method : `POST`
- Required:
  - Body:

    | body  |
    | ------ |
    | Title |
    | Poster |
    | Description |
    
- Request Body:
    ```json
      {
          "Title": "Apartemen Mlati Asri",
          "Poster": "Mlati Asri",
          "Description": "Apartemen hunian keluarga, 3 juta/bulan.",
      }
    ```

- Response :

    ```json
    {
      "ID": 2,
      "CreatedAt": "2025-10-16T19:07:04.319033Z",
      "UpdatedAt": "2025-10-16T19:07:04.319033Z",
      "DeletedAt": null,
      "title": "Apartemen Mlati Asri",
      "poster": "Mlati Asri",
      "description": "Apartemen hunian keluarga, 3 juta/bulan."
    }
    
    ```
   
### Delete House Listing

- Path : `/house/{id}`
- Method : `DELETE`
- Response :

   ```json
       {
        "ID": 0,
        "CreatedAt": "0001-01-01T00:00:00Z",
        "UpdatedAt": "0001-01-01T00:00:00Z",
        "DeletedAt": "2025-10-16T19:51:26.482633803Z",
        "title": "",
        "poster": "",
        "description": ""
      }
   ```

   ### Update House Listing

- Path : `/house/{id}`
- Method : `PUT`
- Required:
  - Optional Body (Minimum 1):

    | body  |
    | ------ |
    | Title |
    | Poster |
    | Description |
    
- Request Body:
    ```json
      {
          "Description": "Apartemen hunian keluarga, 2 juta/bulan.",
      }
    ```

- Response :

    ```json
    {
      "ID": 2,
      "CreatedAt": "2025-10-16T19:07:04.319033Z",
      "UpdatedAt": "2025-10-16T19:07:04.319033Z",
      "DeletedAt": null,
      "title": "Apartemen Mlati Asri",
      "poster": "Mlati Asri",
      "description": "Apartemen hunian keluarga, 2 juta/bulan."
    }
    
    ```

