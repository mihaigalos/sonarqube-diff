# sonarqube-diff 

Diff generated sonarqube reports to detect quality regressions.

### Why?

Quality gates can be set up to guard against regressions detected by sonarqube.
Most likely, accumulated technical debt means one cannot simply turn "hard" gating on, but instead must rely on "no new regressions / no further technical debt" detected by sonarqube.