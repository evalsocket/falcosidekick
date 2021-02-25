package outputs

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"strings"

	"github.com/falcosecurity/falcosidekick/types"
	stan "github.com/nats-io/stan.go"
)

// StanPublish publishes event to NATS Streaming
func (c *Client) StanPublish(falcopayload types.FalcoPayload) {
	c.Stats.Stan.Add(Total, 1)

	nc, err := stan.Connect(c.Config.Stan.ClusterID, c.Config.Stan.ClientID, stan.NatsURL(c.EndpointURL.String()))
	if err != nil {
		c.setStanErrorMetrics()
		log.Info("[ERROR] : STAN - %v\n", err.Error())
		return
	}
	defer nc.Close()

	r := strings.Trim(slugRegularExpression.ReplaceAllString(strings.ToLower(falcopayload.Rule), "_"), "_")
	j, err := json.Marshal(falcopayload)
	if err != nil {
		c.setStanErrorMetrics()
		log.Info("[ERROR] : STAN - %v\n", err.Error())
		return
	}

	err = nc.Publish("falco."+strings.ToLower(falcopayload.Priority.String())+"."+r, j)
	if err != nil {
		c.setStanErrorMetrics()
		log.Info("[ERROR] : STAN - %v\n", err)
		return
	}

	// Setting the success status
	go c.CountMetric(Outputs, 1, []string{"output:stan", "status:ok"})
	c.Stats.Stan.Add(OK, 1)
	c.PromStats.Outputs.With(map[string]string{"destination": "stan", "status": OK}).Inc()
	log.Info("[INFO] : STAN - Publish OK\n")
}

// setStanErrorMetrics set the error stats
func (c *Client) setStanErrorMetrics() {
	go c.CountMetric(Outputs, 1, []string{"output:stan", "status:error"})
	c.Stats.Stan.Add(Error, 1)
	c.PromStats.Outputs.With(map[string]string{"destination": "stan", "status": Error}).Inc()
}
