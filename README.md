Task Details:

Create a web application which takes a website URL as an input and provides general information
about the contents of the page:
- HTML Version
- Page Title
- Headings count by level
- Amount of internal and external links
- Amount of inaccessible links
- If a page contains a login form

Solution:
Create .env file and provide actual parameters of the application. 

Docker commands to run the application are as follows:
```
#docker build -t demoservice .

#docker run --rm -p 8080:8080 demoservice
```

How to test for the page "https://github.com/login"
```
#curl -X GET http://localhost:8080/website-details?url=https%3A%2F%2Fgithub.com%2Flogin

#{"website_url":"https://github.com/login","title":"Sign in to GitHub ┬╖ GitHub","html_version":"Html 5","internal_links":8,"external_links":3,"inaccessible_links":0,"login_page":true}

```