Create display layout profile, store display and windows layout as session, restore it.
For Windows only.

```
# Change your setting from Settings > System > Display
dispsess.exe -save  # Create profile.json
dispsess.exe        # Store display and windows layout as session.json, apply profile.json
dispsess.exe        # Restore session.json
```

Build
------

```
GOOS=windows go build -ldflags -H=windowsgui
```
