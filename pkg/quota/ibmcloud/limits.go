package ibmcloud

// NOTE: Cannot get service quotas/limits via API at the moment.
// https://cloud.ibm.com/docs/vpc?topic=vpc-quotas

func loadLimits() ([]record, error) {
	return []record{
		{
			Service: "is.vpc",
			Name:    "vpc",
			Scope:   "region",
			Value:   10,
		},
		{
			Service: "is.vpc",
			Name:    "subnet",
			Scope:   "vpc",
			Value:   15,
		},
		{
			Service: "is.vpc",
			Name:    "security-group",
			Scope:   "vpc",
			Value:   50,
		},
		{
			Service: "is.load-balancer",
			Name:    "alb",
			Scope:   "region",
			Value:   50,
		},
		{
			Service: "is.instance",
			Name:    "cpu",
			Scope:   "region",
			Value:   200,
		},
		{
			Service: "is.instance",
			Name:    "memory",
			Scope:   "region",
			Value:   200, // GB
		},
		{
			Service: "is.instance",
			Name:    "storage",
			Scope:   "region",
			Value:   18000, // GB
		},
		{
			Service: "is.instance",
			Name:    "floating-ip",
			Scope:   "zone",
			Value:   20,
		},
		{
			Service: "is.volume",
			Name:    "volume",
			Scope:   "region",
			Value:   300,
		},
	}, nil
}
