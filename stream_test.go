package json_test

import (
	"fmt"
	"strings"
	"testing"

	"h12.me/mongodb-json"
)

func TestStream(t *testing.T) {
	q := `
[
    {$match: {
        "impression.created_at": {$exists:1},
        "impression.created_at": {$gte: ISODate("2017-02-24T14:06:08+08:00"), $lt: ISODate("2017-02-24T15:06:08+08:00")},
        "response.created_at": {$exists:1},
        "response.ad.bid_type": "cpi"
    }},
    {$group: {
        _id: {
            imp_hour: {y:2017,m:2,d:24,h:14},
            request: {
                device: {
                    os: {
                        name: "$response.request.device.os.name",
                        version: "$response.request.device.os.version"
                    },
                    model: "$response.request.device.model",
                    geo: {
                        country: "$response.request.device.geo.country"
                    }
                },
                slot: {
                    site: {
                        id: "$response.request.slot.site.id",
                        aff_id: "$response.request.slot.site.aff_id",
                        publisher: {
                            id: "$response.request.slot.site.publisher.id",
                        }
                    },
                    tag: "$response.request.slot.tag"
			    },
                ad: {
                    id: "$response.ad.id",
                    adgroup: {
                        id: "$response.ad.adgroup.id"
                    },
                    campaign: {
                        id: "$response.ad.campaign.id"
                    },
                    advertiser: {
                        id: "$response.ad.advertiser.id"
                    }
                }
            }

        },
        impression: { $sum: 1 }
    }},
    {$out: "imp_by_hour"}
]
`

	d := json.NewDecoder(strings.NewReader(q))
	for d.More() {
		tok, err := d.Token()
		if err != nil {
			t.Fatal(err)
		}
		fmt.Println(tok)
	}
}
