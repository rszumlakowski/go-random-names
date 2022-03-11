package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const adjectives = `
    absolute abstract absurd acidic acoustic adamant adept adequate adorable aero
    aged aggregate aggressive agile agreeable alert alfresco alive alkaline aloof
    alpine ambient ambitious amoeboid amorphous amphibious amused analog ancestral
    ancient angora angry angular annoying anonymous anti anxious apocryphal aqua
    arch ardent arid aromatic arrogant artesian artificial ascendant ascetic ashamed
    ashy astute atomic attractive audacious audio auric auspicious authentic autumn
    average awful awkward bald banoffee beautiful bedraggled benevolent better
    bewildered billowing binary bitter bizarre black blatant blue blushing bodacious
    boiling bold bonafide bored bottomless brainy brash brave breakable breezy brief
    bright brilliant broken brutal bulky bumpy buoyant busy byzantine cajun calm
    canine capable carbide careful cashmere cathartic cautious charming chatty
    cheeky cheerful chemical chiffon chilly chromatic churlish citron civil classic
    clean clear clever cloudy clumsy coaxial cognitive cold colourful combative
    common complex complicit compound concerned concise confused conic contra cool
    cordial courageous covalent crafty cranky crashing crazy creamy creepy creole
    crimson crosstown crowded cruel cubic cuddly curious current cute cyan daily
    damaged damp danger dangling dark dawn dazzling deafening deep defeated defiant
    defunct delicate delicious delightful demi dense depressed determined different
    difficult digital diligent dire dirty disco discrete disparate distant distinct
    disturbed divine dizzy dorsal double doubtful drab dread dry ductile dull eager
    early east easy echoing eigen elaborate elastic elated elder electric elegant
    elite emergent empirical empty enchanting endemic energetic enthused envious
    ephemeral equine equivocal erudite eternal excellent excess excited exo
    expensive express exquisite extended extra exuberant faded faint fair faithful
    fake falling false famous fancy fantastic fast fastidious faustian favourite
    feisty feline ferocious fervent feudal fictional fiendish fierce filthy final
    fine finicky finite first fit flaky flamboyant flannel flat flimsy flippant
    floppy floral fluent fluffy foolish forbidden forgotten fragile fragrant frail
    frantic fraught freezing fresh friendly frightened frigid frisky frivolous
    frosty funny furtive future galvanic gaseous gauche gentle genuine geo giant
    gifted giga gilded glamourous glass gleaming gloomy glorious glum gnomic gold
    good gorgeous graceful gradual gray greasy great greedy green greyscale grieving
    grotesque grotty grouchy grubby grumpy halide hallowed handsome happy haptic
    harsh hasty haughty hazy healthy heavy helpful helpless hidden hilarious hissing
    historic hollow homely honest horrible hot howling huge hulky humid hungry hurt
    hyper icy ideal igneous ill immense imperative imperial impolite important
    impossible impressive impromptu incessant incognito indie indigo infamous
    inferior infernal inner innocent inorganic inquisitive insolent instant integral
    intricate intrinsic inverse ionic iron ironic itchy jealous jittery jolly jovian
    joyous juicy jumbo kafkaesque kind kinetic laconic languid large laser late
    lawful lazy leather lemon lethargic light linear lingering linked liquid little
    lively loitering lonely long loose loud lovely lucky lucrative lunar magenta
    magic magnetic malleable mammoth manic maple marble marvellous massive mauve
    mealy mega melodic melted melting merino meta micro middle mighty mineral mini
    misty mithril modern modest mollified monochrome moot motionless mouldy
    muddy muscular mushy musky myriad mysterious mystic naive nameless narrow nasty
    nasvent national natural naughty near nervous neutral nice noble noisy nominal
    north notmyfirst nutritious nutty obedient obligatory oblivious obnoxious odd
    official old open optimistic orange original ornery outer outrageous palpable
    panicky parallel parched partial patient pending perfect personal pertinent
    pestering petite piccolo pink plain plastic platinum platonic pleasant pliable
    plucky plump plush poised polished polite polyester popular potential poutpout
    powerful pragmatic precious prickly primal primitive primodrial pristine
    profound proud pseudo punk puny pure purple purring putrid quadratic quaint
    quant quantum quick quiet radical ragtag rainbow rancid rapid rapping rare raspy
    rational real red redacted redux refined relaxing relentless relieved reluctant
    remarkable repulsive restive restless reverse rhythmic rich ripe rotten rough
    rowdy royal rural rustic rusty sabretooth safety salmon salty sanguine sassy
    savoury scaly scary screeching scruffy second selfish semantic sentient serious
    shaggy shallow shapely sharp shattered shifty shiny short shrilling shy silent
    silly silver simple sincere sketchy skinny sleepy slimy slippery slow small
    smiling smitten smoggy sneaky sneezy snotty snowy social solar solid solitary
    sonic soothing sore sour south space spammity sparkly sparse spartan special
    spherical spicy splendid spoiled spotless spring squeaking squiggly stale
    standard steampunk steep stereo sticky still stinky stocky stoic stompy storied
    stormy stout strange strong stubby sub subdued subtle successful sudden summer
    sumptuous super superb superior suspicious sweet swift symbiotic synthetic
    taciturn takeaway talented tall tame tangy tart tasteless tasty tatami tedious
    teeny telluric temporary tenacious tender tense terrible thankful third thirsty
    thoughtful thundering tinkling tinny tiny tired tough tricky triple trite
    troubled tubular turbo turbulent twilight ugliest ugly ultimate ultra uncanny
    uncouth uneven unit united unkempt unsightly untenable unusual upset uptight
    urban urgent vacant vast velveteen ventral verbose vestigial vexed viable
    victorious video vigilant vigorous virtual viscose visible vital vivacious
    wailing wandering warm weak weary weathered west wet whining whispering white
    wicked wicker wild winter wispy withered witty wonderful wonky wooden woolly
    worried worrisome wrong xenial yellow young yummy zany zealous zen surly
    benevolent lacklustre formidable indulgent economical gruelling matcha
    tropical granite granola cathoderay deceptive whirling deft cultivated
    celestial clockwork copper grizzly hypnotic ivory jade phantasmal prodigal
    hydraulic pneumatic tranquil persuasive respectable flummoxed quixotic
    fallible medical material essential awestruck awesome curated watershed
    brass alabaster creative illustrious bronze concordant cosmic boreal tiki
    spectral sylvan triassic clay abyssal malachite shambling spumoni vanilla
    chocolate quirky rococo baroque crucial peckish dogmatic inscrutable savvy
    palatable robust apathetic specific smarmy exorbitant dasblinken bigole
    inflected major minor medium bilingual frank newfangled eerie niche
    capricious bohemian robo meteoric peameal techno
		`

const nouns = `
      aardvark abbey ablation ablutions accolade acronym actinide addendum advantage
      adventure affinity affogato aisle albatross alchemy algebra algorithm allegory
      alligator allotrope alloy allusion alpaca amalgam amaranth ampersand amplifier
      analogy anarchy android anecdote animal ant antelope anthem aperture apex apogee
      apparel appendix applause aptitude arcology armadillo artiste ashram asterisk
      asteroid auberge audience aurora automata avatar avocado axiom axon babka baboon
      badger bagpipe baguette baklava balalaika balcony ballad ballista balustrade
      banana banjo banner bannock banquet baobab barbarian barber barge baritone barn
      barnacle barometer barracks barracuda baryon bass bassoon bat batch bayou beagle
      beaker beans bear bee beetle beluga bicycle bidet bijou bilby binoculars biome
      bird bison bistro bloom blossom blunder boar bobcat bodega bog bokeh bonobo
      boson bouquet box brain branch breeze brew brilliant brook bubble buddy budgie
      buffalo bugle bungalow bush butler butterfly buttress buzz cabin caboose
      cacophony cactus caiman cake calculus calipers calliope camel canape canoe
      cantina canto canyon cappuccino capybara caracal caravel cardinal caribou
      carousel casbah cascade casino cassata cassowary castle cat catalyst catapult
      catfish cattle cave centipede century chalet chameleon chamois chaos
      chateau cheese cheetah chemistry cherry chicken chimp chinchilla chinook
      chondrite chortle chronicle chuckle chum chupacabra cipher circus city civet
      clam clarinet clarity clash clavier clown cobbler cobra coconut coffee cog
      comedy comet concertina conch conductor conga consensus consulate cookies coral
      core cornucopia corral cortado cottage couch cougar cow coyote crab crane
      credenza creek critter crocodile croissant crumble crustacean crystal cubit
      cuckoo cucumber culdesac cupboard cupcake curio cuttlefish cyborg cymbal damper
      dance darkness dawn debacle debut deer denizen depanneur desert desperado detour
      dew dialect dialog diamond dichotomy digger diner dinghy dingo dinner dinosaur
      diode divisor dodo dog donkey donut doplhin dormouse dragon dragonfly drama
      dream dregs drum duck dude dugong duke dumpling dune durian dust dwarf dynamo
      dynasty eagle earth echidna eclair eclipse edifice eel elbow electron element
      elephant elf ellipse embassy emerald empanada emu enclosure energy enigma
      enthalpy entropy enzyme epee epic epoch equation escarpment espresso essay
      estate esteem estimate estuary ether ethics euphemism excursion fable factor
      factorial falcon falsetto fantasy farce farewell farthing fauna fauxpas feather
      ferret fiction fiddle field fife filament firbolg fire firefly firmament fish
      flamingo flat flora flounder flower flute flux fly focus fog foil foliage folly
      fondue forest forge formula fortress fossa fracas fractal fraction fragment
      frangipani fritter frog frost fugue function fungus funicular furlong furniture
      fuss futon gaffe galaxy galleon galley garage garb gear gecko geometry gerbil
      gibbon gimmick giraffe glacier glade glitter glossary gluon gnome gnu goat
      goblet golem goose gopher gorilla gourd grail grammar granfondo grass grasshoper
      gravity guitar guppy habanero hacienda hadron haiku hamlet hamster harbinger
      hare harmonica harmony harp hart haze heap heath hedgehog hen hero heron herring
      heuristic hibiscus hill hippo hippodrome hobbit hodgepodge horizon hornet hostel
      hotel hound hubris humour hyena hypothesis iceberg idiom igloo iguana impala
      insect insular integer intrigue island isomer isotope isthmus jackal jaguar
      jargon jay jellyfish journey jukebox jumble jungle kangaroo karaoke karma kayak
      kazoo kea keytar kingfisher kiosk kiwi knight koala kohlrabi kolache komissar
      krill kudu ladybug lake lamington landmark latte lattice laughter leaf legend
      lemming lemur leopard lepton liaison library liger limelight limerick lion
      lizard llama lobster locale locust logarithm logic loon lute lynx lyre lyric
      macaque macaron macaw macchiato macguffin machina maelstrom maestro magnate
      magnolia magpie majordomo mammoth manatee mandala mandolin maneuver mango
      mangrove manifesto manor mansion maracas marimba marmalade marsh mastiff math
      matrix mayfly meadow media meeple melody menagerie mesa mesquite metal
      metropolis microscope milieu millenia millipede mime minotaur mire mixture
      modulus mogul molarity mole molecule mollusk mongoose monk monkey monorail
      montage monument mood moon moose morning mosaic motel moth motif mountain mouse
      mule munro muon mystique mythos nadir nanite narwhal nautilus nebula nemesis
      neuron neutrino neutron newt nexus night nimbus ninja noodles notch novella
      numeral numerator oasis obelisk oboe ocarina ocean ocelot octave octet octopus
      octothorpe odyssey ogre onsen oolite opera orbit orchestra oriole ornament
      ostrich otter overture owl oyster pagoda pajamas palace pangolin pantheon
      panther pantomime pantry paper paprika parable parabola paradise paradox parfait
      parody parrot pastiche pathos patio pavilion peacock pearl pedestal pedigree
      pelican peloton pemmican penguin peninsula pepper petal phase pheasant phoneme
      photon physics piano piccolo pickle picnic pie pig pike pine pingo pinnacle
      piranha pizza plain planet platypus plumeria poem polymer pompom pond poodle
      popcorn porcupine posse possum potato poutine prairie prediction pretzel
      priority prose proton proxy puffin pug puma pumpkin puppet puzzle pyramid quail
      quark quasar quatrain quest rabbit raccoon racket radix ragdoll rain ramen ranch
      rascal rat reaction reagent reality realms reduction reef refrain register
      reindeer relay replicant resonance rhetoric rhino riddle river roadrunner robin
      robot rock rocketship rodent rodeo romance ronin rooster rubbish rubble ruby
      ruffian rumours rumpus rutabaga saga salon saloon samovar samurai sapphire
      sasquatch satellite satire savanna scallop scenario scene schnitzel schooner
      science scorpion scow sea seal sealion sequence shack shadow shaman shanty shape
      shark shed sheep shepherd shoe shore shrimp shrubbery sidecar silence silo
      simile skunk sky sleigh slug smoke snail snapper snow snowflake snowman soba
      sojourn solution solvent somersault sonnet sophist sorcerer souffle sound
      spaniel spectacle sphinx spiral sprawl sprocket squash squid squirrel stampede
      star starfish steppe stevedore sticker stork strait stream strudel student stump
      substance subway succulent suede sun sunrise sunset supper surf surprise swamp
      swan symmetry synapse syntax tachyon taiga tamarin tango tanuki tapir tarn
      tarsier telescope telex tenor tension tensor termite terrain terrier terroir
      tetra theory theremin thesaurus thesis thespian thimble thunder tiger tiramisu
      toad toaster tofu topology toque torte tortoise torus toucan tragedy trail tram
      trebuchet tree trestle tribute tricycle trilobite trinket triumph trombone
      troubadour trouble troupe trousers trumpet tuba tuffet tugboat tundra tunnel
      turkey turnip tutor tuxedo udon ukelele ukulele umbrella unicorn unicycle utopia
      vacation vagabond valkyrie varmint vassal vector vegetable veldt vellum velocity
      velodrome veneer vignette villa village violet violin vocabulary voice void
      volcano vortex vulture wabbit waffle wallaby walrus waltz wanderlust warlock
      warthog wasp water waterfall wave weasel wendigo wetland whale whippet whistle
      wig wildebeest wildflower wind wizard wolf wolverine wombat wood woodchuck
      woodpecker wumpus yak yeti yield yodel yurt zeal zebra zebu zenith zest zigzag
      zinc zinnia zircon zither zodiac zoology nonsense cooper hermit cupoftea
      multipass malarkey lagoon epilogue receptor plottwist dervish terrace quilt
      artifact construct ankh camouflage conservator cyclops feedback liege gauntlet
      illusion sanctuary juggernaut djinn pegasus merfolk hydra sprite dryad druid
      basilisk bramble quagmire zeitgeist decade nutshell kittycat puppydog quandary
      culture tempeh nuance inkling knowledge bazaar nomad sprinkles uncial faerie
      glyph boomerang cocoon wyrm equinox solstice swarm kismet relic sentinel
      silhouette tempest voodoo amulet candelabra ornithopter lagomorph praetor
      flotilla astrolabe lodestone aegis aurochs scarab blizzard centaur fumarole
      brownie cauldron jester cohort magus lumberjack strider sleight simulacrum
      tarpan bauble vertigo tower trattoria tartufo candy cannoli antiquity nova
      pulsar dilemma criteria rigamarole flytrap catoblepas rhapsody gadget widget
      doodad gizmo endeavour atmosphere parliament lexicon mnemonic tautology
	`

type Components struct {
	adjective string
	noun      string
	number    int32
}

func pickWord(in string) string {
	if len(in) == 0 {
		return ""
	}

	words := strings.Fields(in)

	if len(words) == 0 {
		return ""
	}

	word := strings.TrimSpace(words[rand.Intn(len(words))])
	return word
}

func generateComponents(adjectives string, nouns string) Components {
	adjective := pickWord(adjectives)
	noun := pickWord(nouns)
	number := rand.Int31()

	return Components{adjective, noun, number}
}

func (c Components) format() string {
	return fmt.Sprintf("%s-%s-%08x", c.adjective, c.noun, c.number)
}

func generate() string {
	components := generateComponents(adjectives, nouns)
	return components.format()
}

func main() {
	rand.Seed(time.Now().UnixNano())
	fmt.Println(generate())
}
