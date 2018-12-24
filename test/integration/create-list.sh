#!/bin/sh
# dpw@seattle.local
# 2018.01.22
#

list='apples
asparagus
avocado
bananas
beets
bell-peppers
bok-choy
cabbage
carrots
cauliflower
celery
chard
cheese
chick-breast
cucumber
garbanzo-beans
garlic
hummus
kale
kidney-beans
lettuce
milk
olives
olive oil
parmesan-cheese
sweet-potatoes
tomatoes
tomato-sauce
tortillas
turkey balls
water-(zero)
yogurt
zucchini
'

host="http://localhost:9040/api"

for item in $list
do
    item=`echo $item | sed -e 's/-/ /'`

    data='{"title":"'${item}'","category":"produce","status":"open"}'

    echo $data

    curl -v -H 'Content-type: application/json' -d $data $host/list
done

