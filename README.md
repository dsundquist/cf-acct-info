
## Printing Basic Account infromation from Cloudflare

Requires your API Email and Key set in the following environment variables:

```
# for Cloudflare-Go
export CLOUDFLARE_API_EMAIL='your_email@gmail.com'
export CLOUDFLARE_API_KEY='9xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx2'
```

Output looks like: 

```

Username   : your_email@gmail.com
Account ID : 3xxxxxxxxxxxxxxxxxxxxxxxf

| Sites:               | Zone ID:                           | Plan:                | Apex record:       | Proxied
| ------------------   | --------------------------------   | ------------------   | ----------------   | ----
| example1.com         | exxxxxxxxxxxxxxxxxxxxxxxxxxxxxx8   | Business Website     | 45.45.45.76        | true
| example2.com         | fxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx8   | Free Website         | 1.2.3.4            | false
| example3.com         | gxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx8   | Free Website         | 45.45.45.76        | true
| example4.com         | hxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx8   | Business Website     | n/a                | n/a
| example5.com         | ixxxxxxxxxxxxxxxxxxxxxxxxxxxxxx8   | Pro Website          | 1.2.3.4            | true
| example6.com         | jxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx8   | Enterprise Website   | 45.45.45.76        | false
| example7.com         | kxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx8   | Pro Website          | 1.2.3.4            | true
| example8.com         | lxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx8   | Pro Website          | 45.45.45.76        | true
| example9.com         | mxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx8   | Pro Website          | 1.2.3.4            | true
```
