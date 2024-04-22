# c1-baton-formal
Welcome to your new connector!

There are a few operations you are able to do from here:
- `make build` will build the basic connector and check that compilation works.
- `make update-deps` will update all of the local dependencies.
- `make add-dep` will pull down and sync any new dependencies, without updating other dependencies.
- `make lint` will run the linter.


## Sync formal resources

```
BATON_FORMAL_API_KEY=<your formal API key> ./dist/darwin_arm64/c1-baton-formal
```

## Add user to a group

```
BATON_FORMAL_API_KEY=<your formal API key> ./dist/darwin_arm64/c1-baton-formal --grant-entitlement group:<formal_group_id>:member --grant-principal <formal_user_id> --grant-principal-type user
```

## Remove user from a group

```
BATON_FORMAL_API_KEY=<your formal API key> ./dist/darwin_arm64/c1-baton-formal --revoke-grant group:<formal_group_id>:member:user:<formal_user_id>
```