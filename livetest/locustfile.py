from locust import HttpUser, TaskSet, task, between

class UserBehavior(TaskSet):
    @task
    def watch_stream(self):
        headers = {
            "Cookie": "__rg_cookie_id__=be6d1243-5357-4480-a8ca-dd92fcb07fee; _tt_enable_cookie=1; _ttp=1iP4pFekpx9UVVszaZk1dlsQHKL; GCP_IAP_UID=107646932455528360326; cto_bundle=aF7vCl9rczlCcFc4NjJlSExCQUxkcDM5a3NxVzFPQlVIWmlOMzI4VHJXWCUyRkhDYmF3NCUyRldiRyUyQk94dTNHSk9sWjN3a2tZZEdpNHlmJTJCWjF1SUFiUEtVMm5TVnFTSlIlMkZkdUp1Z1Y0MllVOHAycXRTeWElMkI5SjN2YTR6WG9Qc0ZHU2ZPdVRnZnU4aVc4MnMlMkZTNjdPJTJCNEFTbGhSRDRsUTlGOVQ1d1NxcVAyRFIlMkI1NmZhOE4lMkIzS2NwOUE4Nkg0UTFuSFpvQloyYVhzQyUyQk1FMVJNQU1id0klMkZtSXhzcFdhYXNrUTBrJTJCcmxnU2Zpc0xlTGRyZEMzRTFEWVozemFVODlXaEVQS0R6ZEFhcVp1Tkt6NnBCU1R5NEdwUktJMUFKV3ZpOW5ObFdUVjRNVzBjYSUyQm9jTnBkJTJGMnllSUt3ZkIlMkJOOSUyRjhqaVRRWEhwbU56UElLdmJaMEk3cTd3YzdONU1nVHYyWmh1UndlTmpIVHVaWGR3S3ZZJTNE; _ga_PQQLJEY9WY=GS1.1.1711079584.4.0.1711079593.51.0.0; _gid=GA1.2.1084954689.1719733282; isLoggedIn=true; _gcl_au=1.1.547160821.1719810163; _ga_6NZ4VTK1B2=GS1.1.1719810714.1.1.1719810903.0.0.0; role=student; name=Zaqia%20Testing%20Empat%20Satu; profpic=https%3A%2F%2Fimgix3.ruangguru.com%2F%2Fassets%2Finitials-rounded%2FZT_model_1.jpg%3Fw%3D360; userID=ZAQIAO0NIC44W7C3; _ga_WDHNGF8DG4=GS1.1.1719889059.187.1.1719889062.0.0.0; __tracker_session_id__=cdb149ce-64fc-4f72-9b80-6175b29dfcfe; token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJydCI6ImV5SmhiR2NpT2lKSVV6STFOaUlzSW5SNWNDSTZJa3BYVkNKOS5leUpqYVdRaU9pSTNNamxSUjI5T05ERkJJaXdpWlhod0lqb3hOekl6TlRBek56TXlMQ0oxYVdRaU9qRTRNRFkyT0Rrd01Dd2lkVzlqSWpvaVdrRlJTVUZQTUU1SlF6UTBWemRETXlJc0luSWlPaUp6ZEhWa1pXNTBJaXdpWkdsa0lqb2lZbVUyWkRFeU5ETXROVE0xTnkwME5EZ3dMV0U0WTJFdFpHUTVNbVpqWWpBM1ptVmxJaXdpWkc0aU9pSk9iMjVsSWl3aWRHOXJaVzVKUkNJNklqRTNNVGs1TURNMk56STVNakV3T1RneE9UZ2lmUS5ISXk1OU14bHh3WU96SDk2dWpDSmtpTEFWd1ZhcE5YUzg4ZkZjZ1RPaUY0IiwiZXhwIjoxNzIyNDk5MjcyLCJ1aWQiOjE4MDY2ODkwMCwidW9jIjoiWkFRSUFPME5JQzQ0VzdDMyIsInIiOiJzdHVkZW50IiwiZGlkIjoiYmU2ZDEyNDMtNTM1Ny00NDgwLWE4Y2EtZGQ5MmZjYjA3ZmVlIiwiZG4iOiJOb25lIiwidG9rZW5JRCI6IjE3MTk5MDM2NzI5MjEwOTgxOTgifQ.8-rZo2HgrULG0yLMiuRz63mF99v07OsSmNATHY3JQcE; refreshToken=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjaWQiOiI3MjlRR29ONDFBIiwiZXhwIjoxNzIzNTAzNzMyLCJ1aWQiOjE4MDY2ODkwMCwidW9jIjoiWkFRSUFPME5JQzQ0VzdDMyIsInIiOiJzdHVkZW50IiwiZGlkIjoiYmU2ZDEyNDMtNTM1Ny00NDgwLWE4Y2EtZGQ5MmZjYjA3ZmVlIiwiZG4iOiJOb25lIiwidG9rZW5JRCI6IjE3MTk5MDM2NzI5MjEwOTgxOTgifQ.HIy59MxlxwYOzH96ujCJkiLAVwVapNXS88fFcgTOiF4; expireToken=1722499092000; _gat_UA-196723136-1=1; _ga=GA1.1.1115190248.1691662555; _ga_XXZDPTKN3B=GS1.2.1719903674.7.0.1719903674.0.0.0; _clck=1nf9jom%7C2%7Cfn4%7C0%7C1317; _clsk=1poo9on%7C1719903675887%7C1%7C1%7Cr.clarity.ms%2Fcollect; _ga_KGEN8KBRBW=GS1.1.1719903674.8.0.1719903711.0.0.0; _ga_19FFXJ19GZ=GS1.1.1719903674.4.0.1719903711.23.0.0; cf_clearance=QnhuqlWNxdNl9trMrpe70gD__UTepHflN4QotgPumug-1719903712-1.0.1.1-qRJJjcqLRgAxzhFp.X4zFRS2Ys2VhF8crL3k9H9c_yHy1e0d1MjWOfsoBh4IamPiwIDZ4c4tc_5vrjmrooIQOQ",  # Example header
            "User-Agent": "locust-test-agent",
            # Add other headers as needed
        }
        self.client.get("https://live-teaching-ba-staging.sirogu.com/live/94426498-116a-480a-999e-e6500dced4f8", headers=headers)

class WebsiteUser(HttpUser):
    tasks = [UserBehavior]
    wait_time = between(1, 5)  # Simulates users waiting between 1 to 5 seconds between tasks
