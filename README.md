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
docker run --rm -it mihaigalos/sonarkube-diff https://raw.githubusercontent.com/mihaigalos/sonarqube-diff/main/data_example/demo_baseline.html https://raw.githubusercontent.com/mihaigalos/sonarqube-diff/main/data_example/demo_baseline_3additions.html
```

Produces the following CSV output representing 3 newly introduced sonarqube entries:

```bash
AWK40INQ-pl6AHs22Mod-manually-added,modules/jdbc-pool/src/main/java/org/apache/tomcat/jdbc/pool/PoolUtilities.java,26,squid:S2068
AWK40IH6-pl6AHs22Mgc-manually-added,java/org/apache/tomcat/util/net/jsse/PEMFile.java,115,squid:ClassVariableVisibilityCheck
AWK40HIg-pl6AHs22K6U-manually-added,java/javax/transaction/xa/XAException.java,29,squid:ClassVariableVisibilityCheck
```
