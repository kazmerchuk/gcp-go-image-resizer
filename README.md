Resizing and Serving images on Google Cloud Storage
==================================

Utilize the full power of Google Cloud Engine and the [Image Golang API](https://cloud.google.com/appengine/docs/legacy/standard/go111/reference/latest/image) to resize your images.


## Setup

1. Clone the repository:

```
git clone git@github.com:kazmerchuk/gcp-go-image-resizer.git
```

2. Deploy to App Engine.

```
gcloud app deploy
```

## Usage

1. Obtain a serving URL for an existing file on Google Cloud Storage by calling the service with the necessary parameters:

```
curl https://PROJECT_NAME.appspot.com/image-url?bucket=bucket_name&image=image_path.jpg
```

2. This will return a URL similar to:

```
http://lh3.googleusercontent.com/ZVPNrWEe3dCIRF51Dz29kEFONz2cD_BR9-nrsTtmL3nUWgwBwWOYkMkDo-eeUnMYCO1Iumankb2pgquMgSq_0ft-0d3WHFL7vUAQUWnEuoby3h_o8bkmNio
```

You can dynamically modify this image as needed. See the Examples section.

## Google Cloud Storage Setup

Note that you will need to grant **Storage Object Admin access** for your GCS objects to a GAE service account responsible for generating URLs. The service account looks like:

```
your-project-id@appspot.gserviceaccount.com
```

## Examples

* By default it returns an image of a maximum length of 512px. [(link)](http://lh3.googleusercontent.com/ZVPNrWEe3dCIRF51Dz29kEFONz2cD_BR9-nrsTtmL3nUWgwBwWOYkMkDo-eeUnMYCO1Iumankb2pgquMgSq_0ft-0d3WHFL7vUAQUWnEuoby3h_o8bkmNio)

* By appending the =sXX to the end of it where XX can be any integer in the range of 0-1600 and it will result to scale down the image to longest dimension without affecting the original aspect ratio. [(link =s256)](http://lh3.googleusercontent.com/ZVPNrWEe3dCIRF51Dz29kEFONz2cD_BR9-nrsTtmL3nUWgwBwWOYkMkDo-eeUnMYCO1Iumankb2pgquMgSq_0ft-0d3WHFL7vUAQUWnEuoby3h_o8bkmNio=s256)

* By appending =sXX-c a cropped version of that image is being returned as a response. [(link =s400-c)](http://lh3.googleusercontent.com/ZVPNrWEe3dCIRF51Dz29kEFONz2cD_BR9-nrsTtmL3nUWgwBwWOYkMkDo-eeUnMYCO1Iumankb2pgquMgSq_0ft-0d3WHFL7vUAQUWnEuoby3h_o8bkmNio=s400-c)

* By appending =pp-br100-rp-s200 the image is smartly cropped, width 100, height 300, quality 100, format JPG. [(link =w100-h300-c-pp-l100-rj)](http://lh3.googleusercontent.com/ZVPNrWEe3dCIRF51Dz29kEFONz2cD_BR9-nrsTtmL3nUWgwBwWOYkMkDo-eeUnMYCO1Iumankb2pgquMgSq_0ft-0d3WHFL7vUAQUWnEuoby3h_o8bkmNio=w100-h300-c-pp-l100-rj)

* By appending =s0 (s zero) the original image is being returned without any resize or modification. [(link =s0)](http://lh3.googleusercontent.com/ZVPNrWEe3dCIRF51Dz29kEFONz2cD_BR9-nrsTtmL3nUWgwBwWOYkMkDo-eeUnMYCO1Iumankb2pgquMgSq_0ft-0d3WHFL7vUAQUWnEuoby3h_o8bkmNio=s0)


## Advanced Parameters

You can find a list of all available parameters at [Stackoverflow](https://stackoverflow.com/a/25438197/1312280)