# sonarqube-diff 

Diff generated sonarqube reports to detect quality regressions.

### Why?

Quality gates can be set up to guard against regressions detected by sonarqube.
Most likely, accumulated technical debt means one cannot simply turn "hard" gating on, but instead must rely on "no new regressions / no further technical debt" detected by sonarqube.

### Usage

```bash
docker run --rm -it mihaigalos/sonarkube-diff url_A urlB # Point these URLs to correct raw generated sonarqube html files.
```

### Example

```bash
just docker_run https://raw.githubusercontent.com/mihaigalos/sonarqube-diff/main/data_example/demo_baseline.html https://raw.githubusercontent.com/mihaigalos/sonarqube-diff/main/data_example/demo_baseline_3additions.html
```