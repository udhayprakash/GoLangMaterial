#### TIP1 :  Always Close The Response Body

        client := http.DefaultClient
        resp, err := client.Do(req)
        if err != nil {
            return nil, err
        }
        defer resp.Body.Close()


#### TIP2: Always Use a Timeout

        timeout := time.Duration(5 * time.Second)
        client := http.Client{
            Timeout: timeout,
        }
        client.Get(url)

