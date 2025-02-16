# Overview of Architecture

## Time
Feb. 16th, 2025

## Description
High level overview of architecture of the application for the MVP. This is up for review and should be evaluated and iterated on as time moves. See (images)[#images] for supporting materials for the design.

## Design
When it comes to the high level overview of the video sharing application, we must consider multiple different, but equally important, features. Even though this is the highest overview we can see that we will need, at a minimum, two databases, two services, and an ingress service. Though this is just the beginning, it does highlight a few key parts. Lets get into those parts.

**Ingress Service**
We need to find a way to route our traffic to our two main services. Routing will be the most important part here and the highest level of concern with this design at this moment. The key point is traffic must be routed to the proper service given particular predefined conditions. Such a condition would be that the user's session token has expired and thus needs to be routed to Authentication for re-authentication. 

With this in mind, we must ask the question; should we have the ingress service be nginx? internally constructed router? Or levy service mesh items such as Istio? Routing can be relatively straight forward but where do we check the session tokens? Is the ingress routing service just meant to route to the server, check session token, and redirect for authentication? Or does the ingress service do that at the entrypoint? I lean towards the latter however that greatly increases the complexity of our ingress service. That would span the question, does the server even need to check the session token since the ingress already does that? I think it should do that. We could use a service such as OAuth2 Proxy to check for valid session tokens and redirect on our behalf thus making it less complex (ingress service) and doesn't require us to create a manual service. This would leave the server to do it's normal actions without having to check the session token each and every request since that logic would be in OAuth2 Proxy. 

Please review concerns for highlights.

**OAuth2 Proxy and Auth Provider**
Part of the authenication service is storing and checking the user's session token. This is accomplished using OAuth2 Proxy. It will handle the checking and storing, configruation of tokens, etc and will redirect for authenication if token is invalid. This allows for us to  reduce the amount of development overhead while providing a secure session token flow.

We can forward all requests to an Authentication Provider such as Google, Github, Meta, Keycloak, etc. We can use Keycloak for our internal username/password login but that is still up for consideration.

**Server**
The server backend will be provided the session tokens so it can provide items such as profiles. Apart from that, the server handles all the processing, posting, gather, and pushing of items such as videos, comments, profiles, subs, live video, etc. The server will communicate directly with a database (right now CockroachDB as it works best in a HA environment like kubernetes).

## Decision
No decision as of Feb. 16, 2025

## Updates
N/A

## Images
![First revision of the architecture design image that covers the authentication, session token, and backend](./architecture_turtlemen_rev1.png)
![Second revision of the architecture design image that covers the authentication, session token, and backend](./architecture_turtlemen_rev2.png)
