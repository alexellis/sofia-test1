def handle(req):
    """handle a request to the function
    Args:
        req (str): request body
    """

    val = ""
    with open("/var/openfaas/secrets/secret") as f:
        val = f.read()
    
    return "Secret: " + val
