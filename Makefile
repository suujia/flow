
# Generates API code from swagger tool 
.PHONY: gen
gen:
    swagger generate server -t gen -f ./swagger/swagger.yml -A spotter
