# sg

## Local Application Setup 
1. Clone project: `git clone https://github.com/stevenelleman/sg.git`
2. `cd ./sg/app`
3. `yarn start`
4. Navigate to `http://localhost:3000/`
5. Additional information can be found in the [app docs](./services/web-frontend/README.md)

## Docs
- Spec ([./docs/spec.md](./docs/spec.md))
- App Setup Docs ([./services/web-frontend/README.md](./services/web-frontend/README.md))
- Service Deployment ([./deploy/README.md](./deploy/README.md))

Services: 
- Web Frontend Service ([./services/web-frontend/src/components/README.md](./services/web-frontend/src/components/README.md))
- Public API Service ([./services/public-api/README.md](./services/public-api/README.md))

## Code-Amendment Best Practices 
- Don't commit directly to master, make PRs. 
- PRs should be reviewed and be required to have a +1 from non-contributor to the PR.
- PRs should have corresponding tests. 
- PRs should include documentation. Documentation should be kept with its relevant code to minimize drift.
- Correct known trip-wires (little annoyances, mistakes, gotchas) immediately - quick fixes pay off dividends.
- Prefix branch name with creator's initials.

## Tasks 
- [x] Initial Deployment Process [Tuesday]
- [x] Frontend w/ Mock Endpoint/Data [Tuesday]
- [x] API Service w/ Mock Data [Wednesday]
- [x] Postgres DB [Thursday]
- [x] Frontend Linter [Thursday]
- [ ] React-Router [Friday]
- [ ] Redux State Management [Friday]
- [ ] Connect DB to the Frontend [Friday]
- [ ] API Table Migrations 
- [ ] Local VM Setup (Using Tilt preferably)
- [ ] Server-Side Transformations

Backlog:
- [ ] Build Process for Dashboard and API service
- [ ] Microservice Deployment Process
- [ ] API Service Test Pattern (Any way to generalize?)
- [ ] Authz Context and Login
    - [ ] Browser Cookier Manager 
    - [ ] User Info Endpoint
- [ ] Frontend Jest Testing 
- [ ] Frontend Redux Testing
- [ ] Frontend Linter ([eslint](https://www.npmjs.com/package/eslint-config-airbnb))
- [ ] Backend Go Linter
- [ ] API Service Error Handling 
- [ ] API Service Logging
- [ ] Per-Request DB Connection Instantiation
